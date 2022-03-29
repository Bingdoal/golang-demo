package logger

import (
	"go-demo/config"
	"io/ioutil"
	"log"
	"os"
)

var (
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	Debug = log.New(os.Stdout, "[DEBUG] ",
		log.Ldate|log.Ltime|log.Lshortfile)
	if config.Env.GetString("mode") == "prod" {
		Debug.SetOutput(ioutil.Discard)
	}

	Info = log.New(os.Stdout,
		"[INFO] ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,
		"[WARNING] ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(os.Stderr,
		"[ERROR] ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
