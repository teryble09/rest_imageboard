package config

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Password string `json:"password"`
	DBname string `json:"dbname"`
}

func GetConfig() Config {
	data := ReadConfig()

	var config Config
	err := json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("couldn't unmarshal config:\n" + err.Error())
	}

	return config
}

// Возмоможно правильнее было бы взять путь к конфигу из переменной окружения (как у Тулзова)
// или как-нибудь по другому это сделать.
// Было бы интересно посмотреть, какие есть способы 
// и как лучше делать на реальных проектах, но пока не хочется усложнять
func ReadConfig() []byte {
	path, err := os.Getwd()
	if err!=nil {
		log.Fatal("while trying to get os.Getwd(): " + err.Error())
	}

	index := strings.Index(path, "rest_imageboard")
	if index == -1 {
		log.Fatal("current dir doesn't contain 'rest_imageboard'")
	}
	path = path[:index + len("rest_imageboard")]
	path += "/config/config.json"

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("couldn't open file in " + path + " with err: " + err.Error() )
	}
	
	return data
}