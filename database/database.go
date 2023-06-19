package database

import (
	"fmt"

	"github.com/tesla59/shortify/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDB() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("database.db"))
	if err != nil {
		panic("Cannot connect to the database")
	}
	fmt.Println("Connected to the database")
	DBConn.AutoMigrate(&models.URL{})
	fmt.Println("Migration Successful")
}

func GetAllURLs() ([]models.URL, error) {
	var URLs []models.URL
	tx := DBConn.Find(&URLs)
	if tx.Error != nil {
		return []models.URL{}, tx.Error
	}
	return URLs, nil
}

func GetURL(id uint64) (models.URL, error) {
	var URL models.URL
	tx := DBConn.First(&URL, id)
	if tx.Error != nil {
		return URL, tx.Error
	}
	return URL, nil
}

func CreateURL(URL models.URL) error {
	return DBConn.Create(&URL).Error
}

func UpdateURL(URL models.URL) error {
	return DBConn.Save(&URL).Error
}

func DeleteURL(id uint64) error {
	return DBConn.Unscoped().Delete(&models.URL{}, id).Error
}

func FindURLbyShortURL(url string) (models.URL, error) {
	var URL models.URL
	tx := DBConn.Where("short_url = ?", url).First(&URL)
	if tx.Error != nil {
		return models.URL{}, tx.Error
	}
	return URL, nil
}
