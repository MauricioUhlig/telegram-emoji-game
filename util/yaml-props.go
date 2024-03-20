package util

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	// Db struct {
	//      Pg map[string]string `yaml:"postgress"`
	// }
	Telegram struct {
		Key     string
		Chat_Id int `yaml:"chat-id"`
	}
	Request_Url     string `yaml:"request-url"`
	Expected_String string `yaml:"expected-string"`
}

var configs Config

func GetTelegramKey() string {
	return configs.Telegram.Key
}
func GetUrl() string {
	return configs.Request_Url
}
func GetExpectedString() string {
	return configs.Expected_String
}

func FetchYAML() {
	dat, err := os.ReadFile("./configs.yaml")
	if err != nil {
		panic(err.Error())
	}

	err = yaml.Unmarshal([]byte(dat), &configs)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(configs)
}
