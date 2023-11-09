package main

import (
	"fmt"
	"go-project/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	db := initDatabase()
	initRouter(db)
}

func initDatabase() *gorm.DB {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Asia%%2FBangkok",
		"root",
		"P@ssw0rd",
		"localhost",
		"3306",
		"goproject",
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func initRouter(db *gorm.DB) {

	gin := gin.Default()

	handler.InitDefaultHandler(gin)
	handler.InitAuthenticationHandler(db, gin)

	gin.Run()
}
