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

var (
	yellow = color.New(color.FgYellow, color.Underline).SprintFunc()
	red    = color.New(color.FgRed, color.Bold).SprintFunc()
	blue   = color.New(color.FgBlue).SprintFunc()
	pink   = color.New(color.FgHiMagenta, color.Bold).SprintFunc()
)

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

func Info(msg string) {
	err := writeLog("INFO", msg)
	if err != nil {
		color.Red(err.Error())
	}
	fmt.Printf("%s %s", blue("[ INFO ]: ->"), msg)
}

func Warning(msg string) {
	err := writeLog("WARN", msg)
	if err != nil {
		color.Red(err.Error())
	}
	fmt.Printf("%s %s", yellow("[ WARN ]: ->"), msg)
}
