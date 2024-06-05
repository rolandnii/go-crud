package database

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/rolandnii/roland-auth/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB


func Connect(dBUrl string) error {
	db, err := gorm.Open(mysql.Open(dBUrl),&gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	
	if err != nil {
		log.Error(err)

		return err
	}

	
	err = db.AutoMigrate(&model.User{},&model.Company{},&model.Otp{})

	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("Database connected and migrated successfully")
	Db = db

	return nil
}
