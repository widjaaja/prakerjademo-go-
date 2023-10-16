package configs

import (
	"fmt"
	"os"
	"net/url"

	"demo-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

func InitDatabase() {
	// dsn := "postgres://avnadmin:AVNS_h2hsxPs7zF_ljiAf1v4@deploy-postgresql-deploy-postgresql.aivencloud.com:12446/defaultdb?sslmode=require"
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=require",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	conn, _ := url.Parse(dsn)
	conn.RawQuery = "sslmode=verify-ca;sslrootcert=ca.pem"
	
	var err error
	DB, err = gorm.Open(postgres.Open(conn.String()), &gorm.Config{})
	if err != nil {
        panic("failed to connect database")
    }
	InitMigration()
}

func InitMigration(){
	DB.AutoMigrate(&models.Comments{}, &models.User{})
}