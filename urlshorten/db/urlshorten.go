package db

import (
	"fmt"
	"io/ioutil"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	yaml "gopkg.in/yaml.v2"
)

var (
	database     *gorm.DB
	currentValue int
	data         map[string]string
	list         []PathURL
)

// URLShorten is struct contain infor of specific url
type URLShorten struct {
	gorm.Model
	Key   string `gorm:"not null;unique"`
	Value string `gorm:"not null"`
	Count int
}

type PathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

//Init is function create a DB
func Init() error {
	yamlFile, err := ioutil.ReadFile("./db.yaml")
	if err != nil {
		return err
	}
	listURL, err := parseYaml(yamlFile)
	list = listURL
	if err != nil {
		return err
	}
	maps := buildMap(listURL)

	data = maps

	db, err := gorm.Open("mysql", "root:123456@/urlshorten?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	//defer db.Close()
	database = db
	database.AutoMigrate(&URLShorten{})

	return nil

}

func AddURLToYamlFile(key, value string) error {
	data[key] = value
	url := PathURL{Path: key, URL: value}
	list = append(list, url)
	convertYaml, err := yaml.Marshal(&list)
	if err != nil {
		return nil
	}
	err = ioutil.WriteFile("./db.yaml", convertYaml, 0644)
	if err != nil {
		return nil
	}
	return nil
}

func parseYaml(data []byte) ([]PathURL, error) {
	var pathUrls []PathURL
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}

	return pathUrls, nil
}

func buildMap(pathURL []PathURL) map[string]string {
	pathToUrls := make(map[string]string)
	for _, path := range pathURL {
		pathToUrls[path.Path] = path.URL
	}

	return pathToUrls
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

func AllURLYamlFile() (map[string]string, error) {
	file, err := ioutil.ReadFile("./db.yaml")
	if err != nil {
		return nil, err
	}

	urls, err := parseYaml(file)

	if err != nil {
		return nil, err
	}
	result := buildMap(urls)
	return result, nil
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

func changeSruct(key string) []PathURL {
	result := make([]PathURL, len(list)-1)

	for index, value := range list {
		if value.Path != key {
			result[index] = list[index]
		}
	}
	return result

}

func DeteleURLYamlFile(key string) error {
	delete(data, key)
	list = changeSruct(key)
	convertYaml, err := yaml.Marshal(&list)
	if err != nil {
		return nil
	}
	err = ioutil.WriteFile("./db.yaml", convertYaml, 0644)
	if err != nil {
		return nil
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
