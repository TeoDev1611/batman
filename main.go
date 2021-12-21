package main

import (
	"github.com/TeoDev1611/batman/log"
)

/* func main() {
	log.Config.AppName = "test"
	log.Config.FileToLog = "batmantest.log"
	err := log.Init()
	if err != nil {
		panic(err)
	}
	log.Info("an example info")
	log.Warning("an example warning")
	log.Error("an example error")
	log.Fatal("an example fatal")
	print("this will be not printed")
} */

func main() {
	log.LogOpts.Error = "CUSTOM_KEY_ERROR"
	log.LogOpts.Info = "CUSTOM_KEY_INFO"
	log.LogOpts.Warning = "CUSTOM_KEY_WARN"
	log.LogOpts.Fatal = "CUSTOM_KEY_FATAL"

	log.LogOpts.ErrorExit = true
	log.LogOpts.FatalExit = true
}
