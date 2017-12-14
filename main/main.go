package main

import (
	"../logger"
	"runtime"
	"os"
)

func main() {
	for i:=0; i<10; i++ {
		logger.Log.Printf("Server v%s pid=%d started with processes: %d", i, os.Getpid(),runtime.GOMAXPROCS(runtime.NumCPU()))
	}
}
