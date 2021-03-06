package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	database     *gorm.DB
	currentValue int
)

// URLShorten is struct contain infor of specific url
type URLShorten struct {
	gorm.Model
	Key   string `gorm:"not null;unique"`
	Value string `gorm:"not null"`
	Count int
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

// AllURLShorten is func get list url from DB
func AllURLShorten() ([]URLShorten, error) {
	listURL := []URLShorten{}
	err := database.Find(&listURL).Error
	if err != nil {
		return nil, err
	}

	return listURL, nil
}

//DeleteURLShorten is func delete a specific url
func DeleteURLShorten(key string) error {
	listURL := URLShorten{}
	database.Where("key=?", key).First(&listURL)
	err := database.Delete(&listURL).Error
	if err != nil {
		return err
	}
	return nil
}

// Count is func cout how many time url redirect
func Count(value string) error {
	tx := database.Begin()
	tx.Raw("SELECT `count` FROM url_shortens WHERE `value` = ? LIMIT 1 FOR UPDATE").Row().Scan(&currentValue)
	fmt.Println(currentValue)
	currentValue++
	err := tx.Exec("UPDATE url_shortens SET `count` = ? WHERE `value` = ? LIMIT 1", currentValue, value).Error
	if err != nil {
		return nil
	}
	tx.Commit()
	return nil

}
