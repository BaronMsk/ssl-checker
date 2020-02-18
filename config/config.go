package config

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
	"time"
)

type ConfigurationStruct struct {
	Interval time.Duration `yaml:"interval"`
	Timeout time.Duration `yaml:"timeout"`
	Trigger string `yaml:"trigger"`
	Domains []string `yaml:"domains"`
	Notification NotificationTelegramStruct
}

type NotificationTelegramStruct struct {
	Telegram struct{
		Token string `yaml:"token"`
		ChatId int64 `yaml:"chatId"`
	}
}


func NewConfig() ConfigurationStruct {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	config := ConfigurationStruct{}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return config
}
