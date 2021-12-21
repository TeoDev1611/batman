package main

import (
	"github.com/TeoDev1611/batman/log"
)

func main() {
	log.Config.AppName = "test"
	log.Config.FileToLog = "asdasd.log"
	err := log.Init()
	if err != nil {
		panic(err)
	}
	path, _ := log.GetLogPath()
	println(path)
	log.Info("XD")
	log.Warning("asdasd")
}
