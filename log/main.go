package log

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
)

type logData struct {
	Level     string                 `json:"level"`
	Message   string                 `json:"message"`
	Timestamp string                 `json:"timestamp"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
}

type Logger struct {
	fields map[string]interface{}
}

type customLogMessage struct {
	// Custom log key to the file warning level: DEFAULT: WARN
	Warning string
	// Custom log key to the file info level DEFAULT: INFO
	Info string
	// Custom log key to the file error level DEFAULT: ERROR
	Error string
	// Custom log key to the file fatal level DEFAULT: FATAL
	Fatal string
	// Custom log key to the file debug level DEFAULT: DEBUG
	Debug string
	// Exit code with the fatal level
	FatalExit bool
	// Exit code with the error level
	ErrorExit bool
}

// The customlog helper for the customLogMessage struct
var LogOpts = customLogMessage{}

var (
	yellow = color.New(color.FgYellow, color.Underline).SprintFunc()
	red    = color.New(color.FgRed, color.Bold).SprintFunc()
	blue   = color.New(color.FgBlue).SprintFunc()
	pink   = color.New(color.FgHiMagenta, color.Bold).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
)

// Init the struct values
func init() {
	LogOpts.Error = "ERROR"
	LogOpts.Info = "INFO"
	LogOpts.Warning = "WARN"
	LogOpts.Fatal = "FATAL"
	LogOpts.Debug = "DEBUG"
	LogOpts.ErrorExit = false
	LogOpts.FatalExit = true
}

var (
	mu     sync.Mutex
	logChan chan string
	wg     sync.WaitGroup
	once   sync.Once
)

func startWorker() {
	if logChan == nil {
		logChan = make(chan string, Config.BufferSize)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for output := range logChan {
			performWrite(output)
		}
	}()
}

// Close flushes and closes the async log channel
func Close() {
	if Config.Async && logChan != nil {
		close(logChan)
		wg.Wait()
	}
}

func performWrite(output string) {
	mu.Lock()
	defer mu.Unlock()

	path, _ := GetLogPath()
	
	// Check rotation
	if fileInfo, err := os.Stat(path); err == nil {
		if fileInfo.Size()+int64(len(output)) > Config.MaxSize {
			rotate(path)
		}
	}

	file, err2 := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err2 != nil {
		return
	}
	defer file.Close()

	file.WriteString(output)
}

func rotate(path string) {
	// Rotate backups
	for i := Config.MaxBackups - 1; i >= 1; i-- {
		oldPath := fmt.Sprintf("%s.%d", path, i)
		newPath := fmt.Sprintf("%s.%d", path, i+1)
		os.Rename(oldPath, newPath)
	}
	
	// Rename current to .1
	os.Rename(path, path+".1")
}

func writeLog(typelog, msg string, fields map[string]interface{}) error {
	if Config.FileToLog == "default" {
		return errors.New("Fail to get the path you need add the path to log first")
	}

	timeNow := time.Now()
	timeLog := timeNow.Format("2006-01-02 15:04:05")

	data := logData{
		Level:     typelog,
		Message:   msg,
		Timestamp: timeLog,
		Fields:    fields,
	}

	var output string
	if Config.JSONFormat {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return errors.New("Cannot serialize to JSON")
		}
		output = string(jsonData) + "\n"
	} else {
		// Plain text format
		output = fmt.Sprintf("%s %s %s", data.Timestamp, data.Level, data.Message)
		if len(fields) > 0 {
			fieldsData, _ := json.Marshal(fields)
			output += fmt.Sprintf(" fields: %s", string(fieldsData))
		}
		output += "\n"
	}

	if Config.Async {
		once.Do(startWorker)
		logChan <- output
		return nil
	}

	performWrite(output)
	return nil
}

// Global logger with fields
var contextLogger = &Logger{fields: make(map[string]interface{})}

// WithFields returns a new logger with the provided fields
func WithFields(fields map[string]interface{}) *Logger {
	return &Logger{fields: fields}
}

func (l *Logger) Info(msg string) {
	err := writeLog(LogOpts.Info, msg, l.fields)
	if err != nil {
		color.Red(err.Error())
		return
	}
	fmt.Printf("%s %s \n", blue("[ INFO ]: ->"), msg)
}

func (l *Logger) Warning(msg string) {
	err := writeLog(LogOpts.Warning, msg, l.fields)
	if err != nil {
		color.Red(err.Error())
		return
	}
	fmt.Printf("%s %s \n", yellow("[ WARN ]: ->"), msg)
}

func (l *Logger) Error(msg string) {
	err := writeLog(LogOpts.Error, msg, l.fields)
	if err != nil {
		color.Red(err.Error())
		return
	}
	fmt.Printf("%s %s \n", red("[ ERROR ]: ->"), msg)
	if LogOpts.ErrorExit {
		os.Exit(2)
	}
}

func (l *Logger) Fatal(msg string) {
	err := writeLog(LogOpts.Fatal, msg, l.fields)
	if err != nil {
		color.Red(err.Error())
		return
	}
	fmt.Printf("%s %s \n", pink("[ FATAL ]: ->"), msg)
	if LogOpts.FatalExit {
		os.Exit(2)
	}
}

func (l *Logger) Debug(msg string) {
	err := writeLog(LogOpts.Debug, msg, l.fields)
	if err != nil {
		color.Red(err.Error())
		return
	}
	fmt.Printf("%s %s \n", cyan("[ DEBUG ]: ->"), msg)
}

// Global level functions (shorthand for global contextLogger)

// Make a log to the terminal and the file with Info level
func Info(msg string) {
	contextLogger.Info(msg)
}

// Make a log to the terminal and the file with Warning level
func Warning(msg string) {
	contextLogger.Warning(msg)
}

// Make a log to the terminal and the file with Error level
func Error(msg string) {
	contextLogger.Error(msg)
}

// Make a log to the terminal and the file with Fatal level
func Fatal(msg string) {
	contextLogger.Fatal(msg)
}

// Make a log to the terminal and the file with Debug level
func Debug(msg string) {
	contextLogger.Debug(msg)
}
