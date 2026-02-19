package log

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLogging(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "batman-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	Config.AppName = "testapp"
	Config.FileToLog = "test.log"
	Config.FilePathLog = tmpDir

	msg := "This is a test message"
	Info(msg)
	Warning(msg)
	Error(msg)
	Debug(msg)

	logPath := filepath.Join(tmpDir, "test.log")
	content, err := ioutil.ReadFile(logPath)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	if len(lines) != 4 {
		t.Errorf("Expected 4 log lines, got %d", len(lines))
	}

	expectedLevels := []string{LogOpts.Info, LogOpts.Warning, LogOpts.Error, LogOpts.Debug}
	for i, level := range expectedLevels {
		if !strings.Contains(lines[i], level) {
			t.Errorf("Line %d expected to contain %s, got: %s", i, level, lines[i])
		}
		if !strings.Contains(lines[i], msg) {
			t.Errorf("Line %d expected to contain message, got: %s", i, lines[i])
		}
	}
}
