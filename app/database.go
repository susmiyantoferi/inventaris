package app

import (
	"fmt"
	"inventaris/helper"
	"inventaris/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Db() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error load env")
	}

	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbName)

	db, errNew := gorm.Open(mysql.Open(conn), &gorm.Config{})
	helper.PanicErr(errNew)

	err = db.AutoMigrate(&models.Produk{}, &models.Inventaris{})
	if err != nil {
		log.Fatal("AutoMigrate gagal:", err)
	}

	DB = db

	return db

}
