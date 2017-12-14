package logger

import (
	"flag"
	"os"
	"log"
)

var (
	Log          *log.Logger
)

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var logPath = dir + "/log.log"
	flag.Parse()
	var file, createError = os.Create(logPath)
	if createError != nil {
		panic(createError)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logPath)
}
