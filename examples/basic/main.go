package main

import (
	"github.com/TeoDev1611/batman/log"
)

func main() {
	// 1. Basic configuration
	log.Config.AppName = "basic-app"
	log.Config.FileToLog = "app.log"

	// 2. Initialization (creates the log folder and file)
	err := log.Init()
	if err != nil {
		panic(err)
	}

	// 3. Simple logging
	log.Info("This is an information message")
	log.Warning("This is a warning message")
	log.Error("This is an error message")
	log.Debug("This is a debug message")

	// Fatal will call os.Exit(2) by default
	// log.Fatal("This is a fatal message")
}
