package log

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

type logData struct {
	TypeOfLog string
	Message   string
	TimeStamp string
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
)

// Init the struct values
func init() {
	LogOpts.Error = "ERROR"
	LogOpts.Info = "INFO"
	LogOpts.Warning = "WARN"
	LogOpts.Fatal = "FATAL"
	LogOpts.ErrorExit = false
	LogOpts.FatalExit = true
}

func writeLog(typelog, msg string) error {
	if Config.FileToLog == "default" {
		return errors.New("Fail to get the path you need add the path to log first")
	}
	timeLog, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	data := logData{
		TypeOfLog: typelog,
		Message:   msg,
		TimeStamp: timeLog.String(),
	}

	jsondata, err3 := json.Marshal(data)
	if err3 != nil {
		return errors.New("Error in parse from struct to json report github")
	}

	path, _ := GetLogPath()
	file, err2 := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o644)
	if err2 != nil {
		return errors.New("Cannot read the file")
	}
	defer file.Close()

	if _, err := file.WriteString(string(jsondata)); err != nil {
		return errors.New("Cannot write the data to the file")
	}

	return nil
}

// Make a log to the terminal and the file with Info level
func Info(msg string) {
	err := writeLog(LogOpts.Info, msg)
	if err != nil {
		color.Red(err.Error())
		return
	}
	fmt.Printf("%s %s \n", blue("[ INFO ]: ->"), msg)
}

// Make a log to the terminal and the file with Warning level
func Warning(msg string) {
	err := writeLog(LogOpts.Warning, msg)
	if err != nil {
		color.Red(err.Error())
		return
	}
	fmt.Printf("%s %s \n", yellow("[ WARN ]: ->"), msg)
}

// Make a log to the terminal and the file with Error level
func Error(msg string) {
	err := writeLog(LogOpts.Error, msg)
	if err != nil {
		color.Red(err.Error())
		return
	}
	fmt.Printf("%s %s \n", red("[ ERROR ]: ->"), msg)
	if LogOpts.ErrorExit {
		os.Exit(2)
	}
}

// Make a log to the terminal and the file with Fatal level
func Fatal(msg string) {
	err := writeLog(LogOpts.Fatal, msg)
	if err != nil {
		color.Red(err.Error())
		return
	}
	fmt.Printf("%s %s \n", pink("[ FATAL ]: ->"), msg)
	if LogOpts.FatalExit {
		os.Exit(2)
	}
}
