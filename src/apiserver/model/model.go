package model

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"task/database"
)

type Task struct {
	Id        int    `json:"id"`
	Completed bool   `json:"completed"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	Detail    string `json:"detail"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Deadline  string `json:"deadline"`
}

func ListTask(p Params) ([]*Task, error) {
	log.Trace("start list task")
	sqlCmd, err := buildSql(listTask, p)
	if err != nil {
		log.Errorf("build list task sql failed: %v", err)
		return nil, err
	}
	rows, err := database.Get().Query(sqlCmd)
	if err != nil {
		log.Errorf("exec sql %s failed: %v", sqlCmd, err)
		return nil, err
	}
	var tasks []*Task
	for rows.Next() {
		t := &Task{}
		err = rows.Scan(&t.Id, &t.Completed, &t.Title, &t.Type, &t.Detail, &t.StartTime, &t.EndTime, &t.Deadline)
		if err != nil {
			log.Errorf("scan query results failed: %v", err)
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func GetTask(id string) ([]*Task, error) {
	log.Tracef("start get task %s", id)
	sqlCmd, err := buildSql(getTask, Params{"id": id})
	if err != nil {
		log.Errorf("build get task sql failed: %v", err)
		return nil, err
	}
	rows, err := database.Get().Query(sqlCmd)
	if err != nil {
		log.Errorf("exec sql %s failed: %v", sqlCmd, err)
		return nil, err
	}
	var ts []*Task
	for rows.Next() {
		t := &Task{}
		err = rows.Scan(&t.Id, &t.Completed, &t.Title, &t.Type, &t.Detail, &t.StartTime, &t.EndTime, &t.Deadline)
		if err != nil {
			log.Errorf("scan query results failed: %v", err)
			return nil, err
		}
		ts = append(ts, t)
	}
	return ts, nil
}

func UpdateTask(value map[string]interface{}) error {
	log.Tracef("start update task with value: %v", value)
	sqlCmd, err := buildSql(updateTask, value)
	if err != nil {
		log.Errorf("build sql for update task failed: %v", err)
		return err
	}
	_, err = database.Get().Exec(sqlCmd)
	if err != nil {
		log.Errorf("exec sql %s failed: %v", sqlCmd, err)
		return err
	}
	return nil
}

func DeleteTask(id string) error {
	log.Tracef("start delete task %s", id)
	sqlCmd, err := buildSql(deleteTask, Params{"id": id})
	if err != nil {
		log.Errorf("build sql for delete task failed: %v", err)
		return err
	}
	_, err = database.Get().Exec(sqlCmd)
	if err != nil {
		log.Errorf("exec sql %s failed: %v", sqlCmd, err)
		return err
	}
	return nil
}

func CreateTask(p Params) error {
	log.Tracef("start create task: %v", p)
	sqlCmd, err := buildSql(createTask, p)
	if err != nil {
		log.Errorf("build sql for delete task failed: %v", err)
		return err
	}
	_, err = database.Get().Exec(sqlCmd)
	if err != nil {
		log.Errorf("exec sql %s failed: %v", sqlCmd, err)
		return err
	}
	return nil
}

func CreateSubTask(p Params) error {
	log.Tracef("start create sub task: %v", p)
	sqlCmd, err := buildSql(createSubTask, p)
	if err != nil {
		log.Errorf("build sql for create sub task failed: %v", err)
		return err
	}
	_, err = database.Get().Exec(sqlCmd)
	if err != nil {
		log.Errorf("exec sql [%s] failed: %v", sqlCmd, err)
		return err
	}
	return nil
}

type SubTask struct {
	Id        int    `json:"id"`
	TaskId    int    `json:"task_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func ListSubTask(taskId string) ([]*SubTask, error) {
	log.Trace("start list sub task")
	sqlCmd, err := buildSql(listSubTask, Params{"task_id": taskId})
	if err != nil {
		log.Errorf("build list sub task sql failed: %v", err)
		return nil, err
	}
	rows, err := database.Get().Query(sqlCmd)
	if err != nil {
		log.Errorf("exec sql [%s] failed: %v", sqlCmd, err)
		return nil, err
	}
	var subTasks []*SubTask
	for rows.Next() {
		st := &SubTask{}
		err = rows.Scan(&st.Id, &st.TaskId, &st.Title, &st.Completed)
		if err != nil {
			log.Errorf("scan list sub task query results failed: %v", err)
			return nil, err
		}
		subTasks = append(subTasks, st)
	}
	return subTasks, nil
}

func DeleteSubTask(id string) error {
	log.Tracef("start delete sub task %s", id)
	sqlCmd, err := buildSql(deleteSubTask, Params{"id": id})
	if err != nil {
		log.Errorf("build sql for delete sub task failed: %v", err)
		return err
	}
	_, err = database.Get().Exec(sqlCmd)
	if err != nil {
		log.Errorf("exec sql [%s] failed: %v", sqlCmd, err)
		return err
	}
	log.Tracef("exec [%s] success", sqlCmd)
	return nil
}

func UpdateSubTask(p Params) error {
	id, ok := p["id"]
	if !ok {
		msg := fmt.Sprintf("sub task id not specified: %v", p)
		log.Error(msg)
		return fmt.Errorf(msg)
	}
	log.Tracef("start update sub task %s", id)
	sqlCmd, err := buildSql(updateSubTask, p)
	if err != nil {
		log.Errorf("build sql for update sub task failed: %v", err)
		return err
	}
	_, err = database.Get().Exec(sqlCmd)
	if err != nil {
		log.Errorf("exec sql [%s] failed: %v", sqlCmd, err)
		return err
	}
	return nil
}

func Register(p Params) error {
	log.Tracef("start register user: ", p["username"])
	sqlCmd, err := buildSql(register, p)
	if err != nil {
		log.Errorf("build sql for register failed: %v", err)
		return err
	}
	_, err = database.Get().Exec(sqlCmd)
	if err != nil {
		log.Errorf("exec sql [%s] failed: %v", sqlCmd, err)
		return err
	}
	return nil
}

type User struct {
	Id       int
	Name     string
	Password string
}

func GetUserByName(username string) ([]*User, error) {
	log.Tracef("start get user by name: %s", username)
	p := Params{
		"username": username,
	}
	sqlCmd, err := buildSql(getUserByName, p)
	if err != nil {
		log.Errorf("build sql for get user by name failed: %v", err)
		return nil, err
	}
	rows, err := database.Get().Query(sqlCmd)
	if err != nil {
		log.Errorf("exec sql [%s] failed: %v", sqlCmd, err)
		return nil, err
	}
	var users []*User
	for rows.Next() {
		u := &User{}
		err = rows.Scan(&u.Id, &u.Name, &u.Password)
		if err != nil {
			log.Errorf("scan get user by name query results failed: %v", err)
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
