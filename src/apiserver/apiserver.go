package apiserver

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"task/global"
	"time"

	"task/apiserver/controller"
)

func logger(c *gin.Context) {
	start := time.Now()
	c.Next()

	e := log.WithFields(map[string]interface{}{
		"latency": time.Now().Sub(start),
		"client_ip": c.ClientIP(),
		"method": c.Request.Method,
		"path": c.Request.URL.Path,
		"errors": c.Errors.ByType(gin.ErrorTypePrivate).String(),
	})

	e.Info()
}

func Start() error {
	r := gin.New()
	r.Use(logger)
	r.Use(gin.Recovery())


	r.LoadHTMLFiles(global.TaskHTMLPath)

	r.GET("/", controller.Hello)
	r.GET("/task", controller.TaskHome)

	err := r.Run()
	if err != nil {
		log.Errorf("run gin instance failed: %v", err)
		return err
	}

	return nil
}
