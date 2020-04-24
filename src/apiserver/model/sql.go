package model

import (
	"bytes"
	"fmt"
	"text/template"

	"task/util"
)

func init() {
	sqlBuilders = map[sqlBuilderEnum]*sqlBuilder{
		sqlBuilderListTask:   newSqlBuilder(sqlBuilderListTask),
		sqlBuilderGetTask:    newSqlBuilder(sqlBuilderGetTask),
		sqlBuilderUpdateTask: newSqlBuilder(sqlBuilderUpdateTask),
		sqlBuilderDeleteTask: newSqlBuilder(sqlBuilderDeleteTask),
	}
}

type params map[string]interface{}

const (
	errorSqlBuilderNotFound string = "can not find sql builder"
)

type sqlBuilder struct {
	tmpl string
	hook func(string) string
}

func (b *sqlBuilder) build(p params) (string, error) {
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

type sqlBuilderEnum int

const (
	sqlBuilderListTask sqlBuilderEnum = iota
	sqlBuilderGetTask
	sqlBuilderUpdateTask
	sqlBuilderDeleteTask
)

const (
	sqlTmplListTask   = `select * from tasks;`
	sqlTmplGetTask    = `select * from tasks where id={{.id}};`
	sqlTmplUpdateTask = `update tasks set
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
where id = {{.id}};`
	sqlTmplDeleteTask = `delete from tasks where id={{.id}};`
)

var sqlBuilders map[sqlBuilderEnum]*sqlBuilder

func newSqlBuilder(be sqlBuilderEnum) *sqlBuilder {
	var tmpl string
	var hook func(sql string) string
	switch be {
	case sqlBuilderListTask:
		tmpl = sqlTmplListTask
	case sqlBuilderGetTask:
		tmpl = sqlTmplGetTask
	case sqlBuilderUpdateTask:
		tmpl = sqlTmplUpdateTask
		hook = func(sql string) string {
			return util.DeleteLast(sql, ",")
		}
	case sqlBuilderDeleteTask:
		tmpl = sqlTmplDeleteTask
	default:
		return nil
	}
	return &sqlBuilder{
		tmpl: tmpl,
		hook: hook,
	}
}

func buildSql(sbe sqlBuilderEnum, p params) (string, error) {
	sb, ok := sqlBuilders[sbe]
	if !ok {
		return "", fmt.Errorf(errorSqlBuilderNotFound)
	}
	sqlCmd, err := sb.build(p)
	if err != nil {
		return "", err
	}
	return sqlCmd, nil
}
