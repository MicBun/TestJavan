package config

import (
	"github.com/MicBun/TestJavan/models"
	"github.com/MicBun/TestJavan/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	dbConnection := utils.GetEnv("DB_CONNECTION", "mysql")
	var db *gorm.DB
	var err error
	if dbConnection == "postgres" {
		dbUrl := "postgres://myuser:yN5Do2dZnQivCkSqulOK9mwWa6VPB5Dx@dpg-cetjjlpgp3jmgl18q2rg-a/activity_tracking"
		db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	} else if dbConnection == "mysql" {
		username := utils.GetEnv("DB_USERNAME", "myuser")
		password := utils.GetEnv("DB_PASSWORD", "mypassword")
		host := utils.GetEnv("DB_HOST", "tcp(mysql:3306)")
		database := utils.GetEnv("DB_DATABASE", "activity_tracking")
		dsn := username + ":" + password + "@" + host + "/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&models.Family{}, &models.Asset{})
	if err != nil {
		return nil
	}
	return db
}
