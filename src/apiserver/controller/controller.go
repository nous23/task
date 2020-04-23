package controller

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"task/apiserver/model"
	"task/global"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

func TaskHome(c *gin.Context) {
	c.HTML(http.StatusOK, global.TaskHTMLFile, nil)
}

func ListTask(c *gin.Context) {
	tasks, err := model.ListTask()
	if err != nil {
		log.Errorf("list task failed: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}
