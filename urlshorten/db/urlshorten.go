package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var database *gorm.DB

// URLShorten is struct contain infor of specific url
type URLShorten struct {
	gorm.Model
	Key   string `gorm:"not null"`
	Value string `gorm:"not null"`
}

//Init is function create a DB
func Init() error {
	db, err := gorm.Open("mysql", "root:123456@/urlshorten?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	//defer db.Close()
	database = db
	database.AutoMigrate(&URLShorten{})

	return nil

}

// AddURLShorten is func add url to DB
func AddURLShorten(key, value string) error {

	urlshorten := URLShorten{Key: key, Value: value}
	err := database.Create(&urlshorten).Error

	if err != nil {
		return err
	}
	return nil
}

func AllURLShorten() ([]URLShorten, error) {
	listURL := []URLShorten{}
	err := database.Find(&listURL).Error
	if err != nil {
		return nil, err
	}

	return listURL, nil
}

func DeleteURLShorten(key string) error {
	listURL := URLShorten{}
	database.Where("key=?", key).First(&listURL)
	err := database.Delete(&listURL).Error
	if err != nil {
		return err
	}
	return nil
}
