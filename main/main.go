package main

import (
	"../logger"
	"../netutils"
	"../search"
	"runtime"
	"os"
)

func main() {
	ip := "192.168.1."
	number := netutils.IpToNumber(ip)
	logger.Info.Printf("Server %d pid=%d started with processes: %d", number, os.Getpid(),runtime.GOMAXPROCS(runtime.NumCPU()))
    newIp := netutils.NumberToIp(number)
	logger.Error.Printf("Server %s pid=%d started with processes: %d", newIp, os.Getpid(),runtime.GOMAXPROCS(runtime.NumCPU()))
	search.Test_solr()
}
