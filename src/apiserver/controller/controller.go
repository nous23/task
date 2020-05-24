package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"task/apiserver/model"
	"task/global"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

func Task(c *gin.Context) {
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
	tasks, err := model.GetTask(id)
	if err != nil {
		log.Errorf("get task %s failed: %v", id, err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if len(tasks) == 0 {
		msg := fmt.Sprintf("task %s does not exist", id)
		log.Error(msg)
		c.String(http.StatusInternalServerError, msg)
		return
	}
	if len(tasks) > 1 {
		msg := fmt.Sprintf("more than one task with id %s", id)
		log.Error(msg)
		c.String(http.StatusInternalServerError, msg)
		return
	}
	c.JSON(http.StatusOK, tasks[0])
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

func CreateTask(c *gin.Context) {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Errorf("read request body failed: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	p := make(map[string]interface{})
	err = json.Unmarshal(b, &p)
	if err != nil {
		log.Errorf("unmarshal create task request body failed: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	err = model.CreateTask(p)
	if err != nil {
		log.Errorf("create task in db failed: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusCreated)
}

func CreateSubTask(c *gin.Context) {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Errorf("read create subtask request body failed: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	p := make(map[string]interface{})
	err = json.Unmarshal(b, &p)
	if err != nil {
		log.Errorf("unmarshal create subtask request body failed: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	err = model.CreateSubTask(p)
	if err != nil {
		log.Errorf("create subtask in db failed: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusCreated)
}

func ListSubTask(c *gin.Context) {
	taskId := c.Param("task_id")
	subTasks, err := model.ListSubTask(taskId)
	if err != nil {
		log.Errorf("list sub task failed: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, subTasks)
}

func DeleteSubTask(c *gin.Context) {
	id := c.Param("id")
	err := model.DeleteSubTask(id)
	if err != nil {
		log.Errorf("delete sub task %s failed: %v", id, err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func UpdateSubTask(c *gin.Context) {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Errorf("read update sub task request body failed: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	p := model.Params{}
	err = json.Unmarshal(b, &p)
	if err != nil {
		log.Errorf("unmarshal update sub task request body failed: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	err = model.UpdateSubTask(p)
	if err != nil {
		log.Errorf("update sub task failed: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, global.LoginHTMLFile, nil)
}

func Register(c *gin.Context) {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Errorf("read register request body failed: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	p := model.Params{}
	err = json.Unmarshal(b, &p)
	if err != nil {
		log.Errorf("unmarshal update sub task request body failed: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	err = model.Register(p)
	if err != nil {
		log.Errorf("register failed: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusCreated, global.TaskHTMLFile, nil)
}

func Login(c *gin.Context) {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Errorf("read login request body failed: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	log.Infof("=== %s", string(b))
	p := make(map[string]string)
	err = json.Unmarshal(b, &p)
	log.Infof("=== %+v", p)
	username := p["username"]
	password := p["password"]
	users, err := model.GetUserByName(username)
	if err != nil {
		log.Errorf("get user by name failed: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if len(users) == 0 {
		msg := "用户不存在"
		c.String(http.StatusInternalServerError, msg)
		return
	} else if len(users) > 1 {
		msg := "用户重复"
		c.String(http.StatusInternalServerError, msg)
		return
	}

	if password != users[0].Password {
		msg := "密码不正确"
		c.String(http.StatusInternalServerError, msg)
		return
	}
	c.String(http.StatusOK, "login success")
}
