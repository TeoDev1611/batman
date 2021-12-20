package log

import (
	"os"
	"path/filepath"

	"github.com/TeoDev1611/batman/errors"
	"github.com/kirsle/configdir"
)

type config struct {
	AppName     string
	FileToLog   string
	FilePathLog string
}

var Config = config{}

func init() {
	Config.AppName = "batman"
	Config.FileToLog = "batman.log"
	Config.FilePathLog = "default"
}

func initApp() (error, string) {
	ConfigPath := configdir.LocalConfig(Config.AppName)
	err := configdir.MakePath(ConfigPath)
	errCheck := errors.CheckErrors(err, "Error in make the app dir")
	if errCheck != nil {
		return errCheck, ""
	}
	return nil, ConfigPath
}

func CreateLog() error {
	err, configpath := initApp()
	errCheck := errors.CheckErrors(err, "Error in the creation of the log")
	if errCheck != nil {
		return errCheck
	}
	Config.FilePathLog = filepath.Join(configpath, Config.FileToLog)
	fh, err := os.Create(Config.FileToLog)
	errFileCreation := errors.CheckErrors(err, "Error in create the file of logs")
	if errFileCreation != nil {
		return errFileCreation
	}
	defer fh.Close()
	return nil
}
