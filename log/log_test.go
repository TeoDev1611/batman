package log

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLogging(t *testing.T) {
	tmpDir := t.TempDir()
	Config.AppName = "testapp"
	Config.FileToLog = "test.log"
	Config.FilePathLog = tmpDir
	Config.JSONFormat = false

	msg := "This is a test message"
	Info(msg)
	
	logPath := filepath.Join(tmpDir, "test.log")
	content, err := ioutil.ReadFile(logPath)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if !strings.Contains(string(content), LogOpts.Info) || !strings.Contains(string(content), msg) {
		t.Errorf("Plain log format incorrect: %s", string(content))
	}
}

func TestJSONLogging(t *testing.T) {
	tmpDir := t.TempDir()
	Config.AppName = "testapp-json"
	Config.FileToLog = "test.json"
	Config.FilePathLog = tmpDir
	Config.JSONFormat = true

	msg := "JSON test message"
	Info(msg)

	logPath := filepath.Join(tmpDir, "test.json")
	content, err := ioutil.ReadFile(logPath)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	var data logData
	if err := json.Unmarshal(content, &data); err != nil {
		t.Fatalf("Failed to unmarshal JSON log: %v", err)
	}

	if data.Level != LogOpts.Info || data.Message != msg {
		t.Errorf("JSON log data incorrect: %+v", data)
	}
}

func TestWithFields(t *testing.T) {
	tmpDir := t.TempDir()
	Config.AppName = "testapp-fields"
	Config.FileToLog = "fields.json"
	Config.FilePathLog = tmpDir
	Config.JSONFormat = true

	fields := map[string]interface{}{
		"user_id": 123,
		"action":  "login",
	}
	
	WithFields(fields).Info("User logged in")

	logPath := filepath.Join(tmpDir, "fields.json")
	content, err := ioutil.ReadFile(logPath)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	var data logData
	if err := json.Unmarshal(content, &data); err != nil {
		t.Fatalf("Failed to unmarshal JSON log: %v", err)
	}

	if data.Fields["user_id"].(float64) != 123 || data.Fields["action"] != "login" {
		t.Errorf("Fields not correctly logged: %+v", data.Fields)
	}
}

func TestAsyncLogging(t *testing.T) {
	tmpDir := t.TempDir()
	Config.AppName = "testapp-async"
	Config.FileToLog = "test-async.log"
	Config.FilePathLog = tmpDir
	Config.Async = true
	Config.BufferSize = 10

	msg := "Async test message"
	Info(msg)
	
	// Must call Close to flush the logs in async mode
	Close()

	logPath := filepath.Join(tmpDir, "test-async.log")
	content, err := ioutil.ReadFile(logPath)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if !strings.Contains(string(content), msg) {
		t.Errorf("Async log not found in file: %s", string(content))
	}
}

func TestLogRotation(t *testing.T) {
	tmpDir := t.TempDir()
	Config.AppName = "testapp-rotate"
	Config.FileToLog = "rotate.log"
	Config.FilePathLog = tmpDir
	Config.Async = false
	Config.MaxSize = 50 // 50 bytes is enough for 1 log line
	Config.MaxBackups = 1

	msg := "A message longer than 50 bytes so it triggers rotation after some writes."
	Info(msg)
	Info(msg) // Second write should trigger rotation

	logPath := filepath.Join(tmpDir, "rotate.log")
	backupPath := logPath + ".1"

	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		t.Errorf("Backup file %s was not created during rotation", backupPath)
	}
}
