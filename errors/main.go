package errors

import "errors"

func CheckErrors(err error, msg string) error {
	if err != nil {
		return errors.New(msg)
	}
	return nil
}
