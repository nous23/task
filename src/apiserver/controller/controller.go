package controller

import (
	"encoding/json"
	"io/ioutil"
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

func GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := model.GetTask(id)
	if err != nil {
		log.Errorf("get task %s failed: %v", id, err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Errorf("read body failed: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	log.Tracef("update body: %s", string(b))
	update := make(map[string]interface{})
	err = json.Unmarshal(b, &update)
	if err != nil {
		log.Errorf("unmarshal json failed: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	update["id"] = c.Param("id")
	err = model.UpdateTask(update)
	if err != nil {
		log.Errorf("update task %s failed: %v", update["id"], err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := model.DeleteTask(id)
	if err != nil {
		log.Errorf("delete task %s failed: %v", id, err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}
