package directories

import (
	"errors"
	"os"
	"path/filepath"

	errs "github.com/TeoDev1611/batman/errors"
)

// Get the CacheDir joined to the path
func GetDir(app string) (string, error) {
	path, err := os.UserCacheDir()
	err2 := errs.CheckErrors(err, "Error in get the cache dir")
	if err2 != nil {
		return "", errors.New("ERror in get the cache dir")
	}
	return filepath.Join(path, app), nil
}
