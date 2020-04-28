package model

import (
	"bytes"
	"fmt"
	"text/template"

	"task/util"
)

type sqlBuilderEnum int

// operation enum
const (
	listTask sqlBuilderEnum = iota
	getTask
	updateTask
	deleteTask
	createTask
	createSubTask
	listSubTask
	deleteSubTask
	updateSubTask
)

var sqlTemplate = map[sqlBuilderEnum]string{
	listTask: `select * from task;`,
	getTask:  `select * from task where id={{.id}};`,
	updateTask: `update task set
{{ if hasKey . "title" }}
title = '{{.title}},'
{{ end }}
{{ if hasKey . "completed" }}
completed = {{.completed}},
{{end}}
{{if hasKey . "type"}}
task_type = '{{.type}}',
{{end}}
{{if hasKey . "detail"}}
detail = '{{.detail}}',
{{end}}
{{if hasKey . "start_time"}}
start_time = '{{.start_time}}',
{{end}}
{{if hasKey . "end_time"}}
end_time = '{{.end_time}}',
{{end}}
{{if hasKey . "deadline"}}
deadline = '{{.deadline}}',
{{end}}
where id = {{.id}};`,
	deleteTask: `delete from task where id={{.id}};`,
	createTask: `insert into task (completed, title, task_type, detail, start_time, deadline) values
({{.completed}}, '{{.title}}', '{{.type}}', '{{.detail}}', '{{.start_time}}', '{{.deadline}}');`,
	createSubTask: `insert into subtask (task_id, title) values ({{.task_id}}, '{{.title}}');`,
	listSubTask:   `select * from subtask where task_id={{.task_id}} order by id;`,
	deleteSubTask: `delete from subtask where id={{.id}};`,
	updateSubTask: `update subtask set
{{if hasKey . "title"}}
title='{{.title}}',
{{end}}
{{if hasKey . "completed"}}
completed={{.completed}},
{{end}}
where id={{.id}};`,
}

var hooks = map[sqlBuilderEnum]func(string) string{
	updateTask: func(sql string) string {
		return util.DeleteLast(sql, ",")
	},
	updateSubTask: func(sql string) string {
		return util.DeleteLast(sql, ",")
	},
}

var sqlBuilders = map[sqlBuilderEnum]*sqlBuilder{
	listTask:      newSqlBuilder(listTask),
	getTask:       newSqlBuilder(getTask),
	updateTask:    newSqlBuilder(updateTask),
	deleteTask:    newSqlBuilder(deleteTask),
	createTask:    newSqlBuilder(createTask),
	createSubTask: newSqlBuilder(createSubTask),
	listSubTask:   newSqlBuilder(listSubTask),
	deleteSubTask: newSqlBuilder(deleteSubTask),
	updateSubTask: newSqlBuilder(updateSubTask),
}

type Params map[string]interface{}

const (
	errorSqlBuilderNotFound string = "can not find sql builder"
)

type sqlBuilder struct {
	tmpl string
	hook func(string) string
}

func (b *sqlBuilder) build(p Params) (string, error) {
	if p == nil {
		return b.tmpl, nil
	}
	t, err := template.New("sql").Funcs(template.FuncMap{
		"hasKey": util.HasKey,
	}).Parse(b.tmpl)
	if err != nil {
		return "", err
	}
	var buffer bytes.Buffer
	err = t.Execute(&buffer, p)
	if err != nil {
		return "", err
	}
	if b.hook != nil {
		return b.hook(buffer.String()), nil
	}
	return buffer.String(), nil
}

func newSqlBuilder(e sqlBuilderEnum) *sqlBuilder {
	tmpl := sqlTemplate[e]
	hook, ok := hooks[e]
	if !ok {
		hook = nil
	}
	return &sqlBuilder{
		tmpl: tmpl,
		hook: hook,
	}
}

func buildSql(e sqlBuilderEnum, p Params) (string, error) {
	sb, ok := sqlBuilders[e]
	if !ok {
		return "", fmt.Errorf(errorSqlBuilderNotFound)
	}
	sqlCmd, err := sb.build(p)
	if err != nil {
		return "", err
	}
	return sqlCmd, nil
}
