package main

import (
	"context"
	"database/sql"
	"fmt"
	"rename-service-name-here/docs"
	"rename-service-name-here/internal/handler"
	"rename-service-name-here/internal/model"
	"rename-service-name-here/internal/util"
	"time"
	_ "time/tzdata"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func main() {
	util.InitLogger()

	initTimeZone()
	config := loadConfig()
	db := initDatabase(config)
	initRouter(db, config)
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		util.Logger.Error("cannot set timezone", zap.Error(err))
	}

	time.Local = ict
}

func loadConfig() util.Config {
	config, err := util.LoadConfig()
	if err != nil {
		util.Logger.Error("cannot load config", zap.Error(err))
	}

	return config
}

func initDatabase(config util.Config) *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Asia%%2FBangkok",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Database,
	)
	sqlDB, err := sql.Open(config.DB.Driver, dsn)
	if err != nil {
		util.Logger.Fatal("cannot connect to db", zap.Error(err))
	}

	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	db, err := gorm.Open(mysql.New(
		mysql.Config{Conn: sqlDB}),
		&gorm.Config{
			Logger:      &SqlLogger{},
			DryRun:      false,
			PrepareStmt: true,
		})
	if err != nil {
		util.Logger.Fatal("cannot open db", zap.Error(err))
	}

	migration(db)

	return db
}

func migration(db *gorm.DB) {
	if err := db.AutoMigrate(&model.UserModel{}); err != nil {
		util.Logger.Fatal("cannot auto migrate db", zap.Error(err))
	}
}

func initRouter(db *gorm.DB, config util.Config) {
	echo := echo.New()
	validator := validator.New()
	con := loadConfig()
	docs.SwaggerInfo.BasePath = con.Swagger.BasePath

	echo.GET("/swagger/*", echoSwagger.WrapHandler)

	handler.InitDefaultHandler(echo)
	handler.InitUserHandler(db, echo, validator)

	echo.Logger.Fatal(echo.Start(fmt.Sprintf(":%v", config.App.Port)))
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n---\n", sql)
}
