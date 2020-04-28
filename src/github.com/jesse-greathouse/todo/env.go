package todo

import (
	"log"
	"os"
)

type Environment struct {
	DIR         string
	BIN         string
	ETC         string
	PKG         string
	SRC         string
	WEB         string
	PORT        string
	STATIC      string
	SQL_PATH    string
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB_NAME     string
	DB_PORT     string
	DSN         string
}

func (e *Environment) Init() {
	// currently executing directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	e.DIR = e.get("DIR", dir)
	e.BIN = e.get("BIN", e.DIR+"/bin")
	e.ETC = e.get("ETC", e.DIR+"/etc")
	e.PKG = e.get("PKG", e.DIR+"/pkg")
	e.SRC = e.get("SRC", e.DIR+"/src")
	e.WEB = e.get("WEB", e.DIR+"/web")
	e.PORT = e.get("PORT", "80")

	e.DB_USER = e.get("DB_USER", "todo")
	e.DB_PASSWORD = e.get("DB_PASSWORD", "todo")
	e.DB_HOST = e.get("DB_HOST", "127.0.0.1")
	e.DB_NAME = e.get("DB_NAME", "todo")
	e.DB_PORT = e.get("DB_PORT", "3306")
	e.DSN = e.DB_USER + ":" + e.DB_PASSWORD + "@(" + e.DB_HOST + ":" + e.DB_PORT + ")/" + e.DB_NAME
	e.STATIC = e.WEB + "/todo/dist/todo"
	e.SQL_PATH = e.SRC + "/sql"
}

func (e *Environment) get(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
