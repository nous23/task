package controller

import (
	"net/http"

	"task/global"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

func TaskHome(c *gin.Context) {
	c.HTML(http.StatusOK, global.TaskHTMLFile, nil)
}
