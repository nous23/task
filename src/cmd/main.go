package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"task/pkg/config"
	"task/pkg/util"
)

func init() {
	initLog()
}

func initLog() {
	if !util.Exist(config.AppConfigs.LogPath) {
		err := os.MkdirAll(filepath.Dir(config.AppConfigs.LogPath), os.ModePerm)
		if err != nil {
			log.Panicf("mkdir %s failed: %v", filepath.Dir(config.AppConfigs.LogPath), err)
		}
		_, err = os.Create(config.AppConfigs.LogPath)
		if err != nil {
			log.Panicf("create log file %s failed: %v", config.AppConfigs.LogPath, err)
		}
	}

	logFile, err := os.OpenFile(config.AppConfigs.LogPath, os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Panicf("open log file %s failed: %v", config.AppConfigs.LogPath, err)
	}
	logrus.SetOutput(logFile)
	logrus.SetFormatter(&logrus.TextFormatter{})
}

func main() {
	logrus.Infof("task start...")
	logrus.Infof("%+v", config.AppConfigs)
}
