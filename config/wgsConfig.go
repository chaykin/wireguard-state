package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type WgsConfig struct {
	WgUpCommand    string `yaml:"wgUpCommand"`
	WgStateCommand string `yaml:"wgStateCommand"`
	WgDownCommand  string `yaml:"wgDownCommand"`
	RxCommand      string `yaml:"rxCommand"`
	TxCommand      string `yaml:"txCommand"`
	RefreshRate    int64  `yaml:"refreshRate"`
}

func ReadConfig(fileName string) *WgsConfig {
	conf := &WgsConfig{}

	file, err := os.OpenFile(fileName, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal("Error opening config file", err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading config file", err)
	}

	if err := yaml.Unmarshal(data, conf); err != nil {
		log.Fatal("Error unmarshalling config file", err)
	}

	return conf
}
