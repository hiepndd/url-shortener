package db

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

var (
	data map[string]string
	list []PathURL
)

//PathURL is struct contain infor of specific url
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
	return nil

}

// AddURLToYamlFile is func add a url to file yaml
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

//AllURLYamlFile is func list all url from file yaml
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

func changeSruct(key string) []PathURL {
	result := make([]PathURL, len(list)-1)

	for index, value := range list {
		if value.Path != key {
			result[index] = list[index]
		}
	}
	return result

}

//DeteleURLYamlFile is func delete a url from file yaml
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
