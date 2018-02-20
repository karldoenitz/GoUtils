package logger

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace        *log.Logger
	Info         *log.Logger
	Warning      *log.Logger
	Error        *log.Logger
)

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var logPath = dir + "/log.log"
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file: ", err)
	}
	Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
