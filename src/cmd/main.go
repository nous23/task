package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"

	"task/apiserver"
	"task/config"
	"task/database"
	"task/util"
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
	formatter := &logrus.TextFormatter{
		PadLevelText: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return frame.Function, fmt.Sprintf("%s:%d", frame.File, frame.Line)
		},
	}
	logrus.SetFormatter(formatter)
	logrus.SetReportCaller(true)
	logrus.SetOutput(logFile)
}

func main() {
	logrus.Infof("task start...")
	logrus.Infof("%+v", config.AppConfigs)

	var err error
	err = database.InitDB()
	if err != nil {
		logrus.Errorf("init db failed: %v", err)
		os.Exit(1)
	}

	err = apiserver.Start()
	if err != nil {
		logrus.Errorf("start apiserver failed: %v", err)
		os.Exit(1)
	}
}
