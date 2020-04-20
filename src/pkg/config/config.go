package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"task/pkg/global"
)

type AppConfig struct {
	LogPath string `json:"log_path"`
}

var AppConfigs AppConfig

func init() {
	b, err := ioutil.ReadFile(global.AppConfigPath)
	if err != nil {
		log.Panicf("read %s failed: %v", global.AppConfigPath, err)
	}

	if err = json.Unmarshal(b, &AppConfigs); err != nil {
		log.Panicf("unmarshal app config failed: %v", err)
	}
}
