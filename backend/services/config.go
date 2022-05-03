package services

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type App struct {
	Port   string
	DBConf *DBConfig
	DB     *sqlx.DB
}

type DBConfig struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}

func Init() *App {

	var a = App{
		Port: "3010",
		DBConf: &DBConfig{
			Host:     "localhost",
			Port:     "3306",
			Name:     "ums_app",
			Username: "root",
			Password: "root",
		},
	}

	db, err := sqlx.Connect("mysql", a.DBConf.getDBConnectionParams())
	if err != nil {
		log.Fatalln(err)
	}

	a.DB = db

	return &a
}

func (d *DBConfig) getDBConnectionParams() string {
	return d.Username + ":" + d.Password + "@(" + d.Host + ":" + d.Port + ")/" + d.Name
}

func (a App) GetDBConn() *sqlx.DB {
	return a.DB
}
