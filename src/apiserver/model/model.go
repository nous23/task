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
	sqlCmd, err := buildSql(sqlBuilderListTask, nil)
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

func GetTask(id string) (*Task, error) {
	log.Tracef("start get task %s", id)
	sqlCmd, err := buildSql(sqlBuilderGetTask, params{"id": id})
	if err != nil {
		log.Errorf("build get task sql failed: %v", err)
		return nil, err
	}
	rows, err := database.Get().Query(sqlCmd)
	if err != nil {
		log.Errorf("exec sql %s failed: %v", sqlCmd, err)
		return nil, err
	}
	t := &Task{}
	for rows.Next() {
		err = rows.Scan(&t.Id, &t.Completed, &t.Title, &t.Type, &t.Detail, &t.StartTime, &t.EndTime, &t.Deadline)
		if err != nil {
			log.Errorf("scan query results failed: %v", err)
			return nil, err
		}
	}
	return t, nil
}

func UpdateTask(value map[string]interface{}) error {
	log.Tracef("start update task with value: %v", value)
	sqlCmd, err := buildSql(sqlBuilderUpdateTask, value)
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
	sqlCmd, err := buildSql(sqlBuilderDeleteTask, params{"id": id})
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
