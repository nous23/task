package apiserver

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"task/apiserver/controller"
	"task/global"
)

func logger(c *gin.Context) {
	start := time.Now()
	c.Next()

	e := log.WithFields(log.Fields{
		"latency":   time.Now().Sub(start),
		"client_ip": c.ClientIP(),
		"method":    c.Request.Method,
		"path":      c.Request.URL.Path,
		"errors":    c.Errors.ByType(gin.ErrorTypePrivate).String(),
	})

	e.Info()
}

func Run() error {
	r := gin.New()
	//r.Use(logger)
	r.Use(gin.Recovery())

	r.LoadHTMLFiles(global.TaskHTMLPath)
	r.LoadHTMLFiles(global.LoginHTMLPath)
	r.Static("/static", global.StaticDir)

	r.GET("/", controller.Hello)

	r.GET("/home", controller.LoginPage)
	r.POST("/register", controller.Register)
	r.GET("/login", controller.Login)

	// task
	r.GET("/task", controller.Task)
	r.GET("/tasks", controller.ListTask)
	r.GET("/task/:id", controller.GetTask)
	r.PUT("/task/:id", controller.UpdateTask)
	r.DELETE("/task/:id", controller.DeleteTask)
	r.POST("/task", controller.CreateTask)

	// sub task
	r.POST("/sub_task", controller.CreateSubTask)
	r.GET("/sub_task/:task_id", controller.ListSubTask)
	r.DELETE("/sub_task/:id", controller.DeleteSubTask)
	r.PUT("/sub_task/:id", controller.UpdateSubTask)

	err := r.Run()
	if err != nil {
		log.Errorf("run gin instance failed: %v", err)
		return err
	}

	return nil
}
