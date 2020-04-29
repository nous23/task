package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"task/apiserver"
	_ "task/log"
)

func main() {
	log.Info("========== task start ==========")

	err := apiserver.Run()
	if err != nil {
		log.Errorf("start apiserver failed: %v", err)
		os.Exit(1)
	}
}
