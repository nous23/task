package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"task/apiserver/controller"
)

func Start() error {
	r := gin.Default()

	r.GET("/", controller.Hello)

	err := r.Run()
	if err != nil {
		logrus.Errorf("run gin instance failed: %v", err)
		return err
	}

	return nil
}
