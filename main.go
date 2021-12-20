package main

import (
	"github.com/TeoDev1611/batman/log"
)

func main() {
	log.Config.AppName = "test"
	log.Config.FileToLog = "asdasd.log"
	err := log.CreateLog()
	if err != nil {
		panic(err)
	}
	println(log.Config.FilePathLog)
	log.Info("asdasd")
}
