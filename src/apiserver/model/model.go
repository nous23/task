package model

import (
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

func ListTask() ([]*Task, error) {
	log.Trace("start list task")
	sqlCmd, err := buildSql(listTask, nil)
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
	sqlCmd, err := buildSql(getTask, params{"id": id})
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
	sqlCmd, err := buildSql(deleteTask, params{"id": id})
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

func CreateTask(p params) error {
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

func CreateSubTask(p params) error {
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
	sqlCmd, err := buildSql(listSubTask, params{"task_id": taskId})
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
	sqlCmd, err := buildSql(deleteSubTask, params{"id": id})
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
