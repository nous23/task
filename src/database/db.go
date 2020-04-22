package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"

	"task/config"
)

var db *sql.DB

func InitDB() error {
	dataSourceName := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",
		config.AppConfigs.DBConfig.User,
		config.AppConfigs.DBConfig.Password,
		config.AppConfigs.DBConfig.Protocol,
		config.AppConfigs.DBConfig.Host,
		config.AppConfigs.DBConfig.Port,
		config.AppConfigs.DBConfig.DBName)
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		logrus.Errorf("open db failed: %v\n", err)
		return err
	}
	return db.Ping()
}
