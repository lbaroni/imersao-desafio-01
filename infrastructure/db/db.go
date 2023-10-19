package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "gorm.io/driver/sqlite"
	
	"github.com/lbaroni/imersao-desafio-01/domain/model"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.env")

	if err != nil {
		log.Fatalf("error loading .env files")
	}

}

func ConnectDB(env string) *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	if env != "test"{
		dsn = os.Getenv("dsn")
		db,err = gorm.Open(os.Getenv("dbType"), dsn)
	} else {
		dsn = os.Getenv("dsnTest")
		db,err = gorm.Open(os.Getenv("dbTypeTest"), dsn)
	}

	if err != nil{
		log.Fatalf("error connecting o database: %v", err)
	}

	if os.Getenv("debug") == "true" {
		db.LogMode(true)
	}

	if os.Getenv("AutoMigrateDb") == "true" {
		log.Printf("Automigrating DB")
		db.AutoMigrate(&model.Product{})
	}

	return db
}