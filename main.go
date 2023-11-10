package main

import (
	"database/sql"
	"fmt"
	"go-project/handler"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	initTimeZone()
	db := initDatabase()
	initRouter(db)
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Fatal("cannot set timezone", err)
	}

	time.Local = ict
}

func initDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Asia%%2FBangkok",
		"root",
		"root",
		"localhost",
		"3307",
		"mysqldb",
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

	return db
}

func initRouter(db *gorm.DB) {

	gin := gin.Default()

	handler.InitDefaultHandler(gin)
	handler.InitAuthenticationHandler(db, gin)

	gin.Run()
}
