package main

import (
	"../logger"
	"../netutils"
	"runtime"
	"os"
)

func main() {
	ip := "192.168.1.169"
	number := netutils.IpToNumber(ip)
	logger.Log.Printf("Server %s pid=%d started with processes: %d", number, os.Getpid(),runtime.GOMAXPROCS(runtime.NumCPU()))
    newIp := netutils.NumberToIp(number)
	logger.Log.Printf("Server %s pid=%d started with processes: %d", newIp, os.Getpid(),runtime.GOMAXPROCS(runtime.NumCPU()))
}
