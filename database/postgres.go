package database

import (
	"fmt"
	"log"
	"os"
	"todo-app_fp1/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
type databaseConfig struct {
    host     string
    port     string
    user     string
    password string
    dbName   string
}

var ginMode = os.Getenv("GIN_MODE")

func init() {
    if ginMode != "release" {
        if err := godotenv.Load(); err != nil {
            log.Fatalln(err.Error())
        }
    }
}

func GetDBConfig() gorm.Dialector {
    dbConfig := databaseConfig{
        host:     os.Getenv("DB_HOST"),
        port:     os.Getenv("DB_PORT"),
        user:     os.Getenv("DB_USER"),
        password: os.Getenv("DB_PASSWORD"),
        dbName:   os.Getenv("DB_NAME"),
    }

    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
        dbConfig.host,
        dbConfig.port,
        dbConfig.user,
        dbConfig.password,
        dbConfig.dbName,
    )

    return postgres.Open(dsn)
}

var db *gorm.DB

func init() {
    var err error
    db, err = gorm.Open(GetDBConfig())
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

    if err := db.AutoMigrate(&entity.Todo{}); err != nil {
        log.Fatalf("Error automigrating database: %v", err)
    }

    log.Println("Connected to DB!")
}

func GetDBInstance() *gorm.DB {
    return db
}