package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"

	"task/config"
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
	formatter := &log.TextFormatter{
		PadLevelText: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return frame.Function, fmt.Sprintf("%s:%d", frame.File, frame.Line)
		},
	}
	log.SetFormatter(formatter)
	log.SetReportCaller(true)
	log.SetOutput(logFile)
}
