package log

import (
	errGo "errors"
	"os"
	"path/filepath"

	"github.com/TeoDev1611/batman/directories"
	"github.com/TeoDev1611/batman/errors"
)

type config struct {
	// Config the basic app setup
	AppName string
	// Filename to log
	FileToLog string
	// Folder to log dont touch :)
	FilePathLog string
}

// Init the struct of config and make public
var Config = config{}

// Init the default values for the struct
func init() {
	Config.AppName = "batman"
	Config.FileToLog = "batman.log"
	Config.FilePathLog = "default"
}

// Make the private function for init the folders for the app
func initFolders() (string, error) {
	path, err := directories.GetDir(Config.AppName)
	errCheck := errors.CheckErrors(err, "Error in get the directory")
	Config.FilePathLog = path
	if errCheck != nil {
		return "", errCheck
	}
	if _, err := os.Stat(Config.FilePathLog); os.IsNotExist(err) {
		err := os.Mkdir(Config.FilePathLog, 0o755)
		if err != nil {
			panic(err)
		}
	}
	return path, nil
}

// Make the util function for get the log file path
func GetLogPath() (string, error) {
	if Config.FilePathLog == "default" {
		return "", errGo.New("The path is undefined error")
	}
	return filepath.Join(Config.FilePathLog, Config.FileToLog), nil
}

// Start the app creating the base empty log file configurated
func Init() error {
	configpath, err := initFolders()
	errCheck := errors.CheckErrors(err, "Error in the creation of the log")
	if errCheck != nil {
		return errCheck
	}
	file := filepath.Join(configpath, Config.FileToLog)
	fh, err := os.Create(file)
	errFileCreation := errors.CheckErrors(err, "Error in create the file of logs")
	if errFileCreation != nil {
		return errFileCreation
	}
	defer fh.Close()
	return nil
}
