package main

import (
	"time"

	"github.com/TeoDev1611/batman/log"
)

func main() {
	// 1. Advanced configuration
	log.Config.AppName = "advanced-app"
	log.Config.FileToLog = "app.json"

	// Enable JSON format (v2 feature)
	log.Config.JSONFormat = true

	// Enable Async logging for high performance (v2 feature)
	log.Config.Async = true
	log.Config.BufferSize = 50 // Buffer for 50 log messages

	// Configure Log Rotation (v2 feature)
	log.Config.MaxSize = 500 * 1024 // 500 KB limit per file
	log.Config.MaxBackups = 5      // Keep up to 5 old log files

	// 2. Initialization
	if err := log.Init(); err != nil {
		panic(err)
	}

	// 3. IMPORTANT: Always close the logger in Async mode to flush logs
	defer log.Close()

	// 4. Logging with Fields (Context support)
	log.WithFields(map[string]interface{}{
		"user_id": 456,
		"action":  "upload",
		"status":  "success",
	}).Info("File uploaded successfully")

	// 5. Standard logging still works with global context
	log.Info("Service is running in async mode")
	log.Debug("Debugging performance metrics")

	// Let's log some more to demonstrate async behavior
	for i := 0; i < 5; i++ {
		log.WithFields(map[string]interface{}{"iteration": i}).Info("Processing data...")
		time.Sleep(10 * time.Millisecond)
	}

	log.Warning("Simulation finished")
}
