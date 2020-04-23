package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"task/apiserver"
	"task/database"
	_ "task/log"
)

func main() {
	log.Info("========== task start ==========")

	var err error
	err = database.InitDB()
	if err != nil {
		log.Errorf("init db failed: %v", err)
		os.Exit(1)
	}

	err = apiserver.Run()
	if err != nil {
		log.Errorf("start apiserver failed: %v", err)
		os.Exit(1)
	}
}
