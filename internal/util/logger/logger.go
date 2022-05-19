package logger

import (
	"go-demo/config"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func InitLogger() {
	level := strings.ToLower(config.Env.GetString("logger"))
	levelMap := map[string]int{
		"debug":   4,
		"info":    3,
		"warning": 2,
		"error":   1,
	}
	levelValue := levelMap[level]

	Debug = log.New(os.Stdout, "[DEBUG] ",
		log.Ldate|log.Ltime|log.Lshortfile)
	if levelValue < levelMap["debug"] {
		Debug.SetOutput(ioutil.Discard)
	}

	Info = log.New(os.Stdout,
		"[INFO] ",
		log.Ldate|log.Ltime|log.Lshortfile)
	if levelValue < levelMap["info"] {
		Info.SetOutput(ioutil.Discard)
	}

	Warning = log.New(os.Stdout,
		"[WARNING] ",
		log.Ldate|log.Ltime|log.Lshortfile)
	if levelValue < levelMap["warning"] {
		Warning.SetOutput(ioutil.Discard)
	}

	Error = log.New(os.Stderr,
		"[ERROR] ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
