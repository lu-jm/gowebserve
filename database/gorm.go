package database

import (
	"fmt"
	"log"
	"test/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var GormDB *gorm.DB

func InitGorm() error {
	dsn := "root:123456@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	GormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		return fmt.Errorf("gorm conntect fail%v", err)
	}
	log.Println("gorm connect success")
	return nil
}

func InitGormWithConfig() error {
	config, err := config.LoadConfig()
	if err != nil {
		return err
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Name)
	// var err error
	GormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		return fmt.Errorf("gorm conntect fail%v", err)
	}
	log.Println("gorm connect success")
	return nil
}

func GetDB() *gorm.DB {
	return GormDB
}
