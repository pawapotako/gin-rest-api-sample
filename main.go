package main

import (
	"database/sql"
	"fmt"
	"go-project/handler"
	"go-project/model"
	"go-project/util"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	initTimeZone()
	config := util.LoadConfig()
	db := initDatabase(config)
	initRouter(db, config)
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Fatal("cannot set timezone", err)
	}

	time.Local = ict
}

func initDatabase(config util.Config) *gorm.DB {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Asia%%2FBangkok",
		config.DbUsername,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbDatabase,
	)
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("cannot connect to db ", err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot open db ", err)
	}

	migration(db)

	return db
}

func migration(db *gorm.DB) {

	if err := db.AutoMigrate(&model.UserModel{}); err != nil {
		log.Fatal("cannot auto migrate db", err)
	}
}

func initRouter(db *gorm.DB, config util.Config) {

	gin := gin.Default()
	validator := validator.New()

	handler.InitDefaultHandler(gin)
	handler.InitAuthenticationHandler(db, gin, validator)

	gin.Run(":" + config.AppPort)
}
