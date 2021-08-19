package config

import (
	"fmt"
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	_ "github.com/joho/godotenv/autoload"
)

func ConnectDB() (c *gorm.DB, err error) {
	DB_CONNECTION := os.Getenv("DB_CONNECTION_ODS")
	DB_HOST := os.Getenv("DB_HOST_ODS")
	DB_PORT := os.Getenv("DB_PORT_ODS")
	DB_DATABASE := os.Getenv("DB_DATABASE_ODS")
	DB_USERNAME := os.Getenv("DB_USERNAME_ODS")
	DB_PASSWORD := os.Getenv("DB_PASSWORD_ODS")

	DB_TEST := os.Getenv("DB_TEST")
	DB_DETAIL := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_DATABASE + "?parseTime=true"
	if DB_CONNECTION == "" {
		DB_DETAIL = DB_TEST
		conn, err := gorm.Open(sqlite.Open(DB_DETAIL), &gorm.Config{})
		if err != nil || conn == nil {
			fmt.Println("Error connecting to DB")
			fmt.Println(err.Error())
		}
		return conn, err
	} else {
		conn, err := gorm.Open(mysql.Open(DB_DETAIL), &gorm.Config{})
		if err != nil || conn == nil {
			fmt.Println("Error connecting to DB")
			fmt.Println(err.Error())
		}
		return conn, err
	}
}

func ConnectDBSY() (c *gorm.DB, err error) {
	DB_CONNECTION := os.Getenv("DB_CONNECTION_SY")
	DB_HOST := os.Getenv("DB_HOST_SY")
	DB_PORT := os.Getenv("DB_PORT_SY")
	DB_DATABASE := os.Getenv("DB_DATABASE_SY")
	DB_USERNAME := os.Getenv("DB_USERNAME_SY")
	DB_PASSWORD := os.Getenv("DB_PASSWORD_SY")

	DB_TEST := os.Getenv("DB_TEST")
	DB_DETAIL := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_DATABASE + "?parseTime=true"
	if DB_CONNECTION == "" {
		DB_DETAIL = DB_TEST
		conn, err := gorm.Open(sqlite.Open(DB_DETAIL), &gorm.Config{})
		if err != nil || conn == nil {
			fmt.Println("Error connecting to DB")
			fmt.Println(err.Error())
		}
		return conn, err
	} else {
		conn, err := gorm.Open(mysql.Open(DB_DETAIL), &gorm.Config{})
		if err != nil || conn == nil {
			fmt.Println("Error connecting to DB")
			fmt.Println(err.Error())
		}
		return conn, err
	}
}
