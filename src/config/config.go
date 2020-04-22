package config

import (
	"bytes"
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v2"

	"task/global"
)

type AppConfig struct {
	LogPath  string   `yaml:"log_path"`
	DBConfig dbConfig `yaml:"db_config"`
}

type dbConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

var AppConfigs AppConfig

func init() {
	b, err := ioutil.ReadFile(global.AppConfigPath)
	if err != nil {
		log.Panicf("read %s failed: %v", global.AppConfigPath, err)
	}
	if err = yaml.Unmarshal(refreshAppConfig(b), &AppConfigs); err != nil {
		log.Println(string(refreshAppConfig(b)))
		log.Panicf("unmarshal app config failed: %v", err)
	}
}

func refreshAppConfig(data []byte) []byte {
	text := string(data)
	u, err := user.Current()
	if err != nil {
		log.Panicf("get current user failed: %v", err)
	}

	m := map[string]string{
		"user_home_dir": filepath.Clean(u.HomeDir),
	}

	tmpl, err := template.New("var").Parse(text)
	if err != nil {
		log.Panicf("parse template %s failed: %v", text, err)
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, m)
	if err != nil {
		log.Panicf("exec template failed: %v", err)
	}
	return b.Bytes()
}
