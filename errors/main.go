package errors

import "errors"

// Check the errors returning the error message util
func CheckErrors(err error, msg string) error {
	if err != nil {
		return errors.New(msg)
	}
	return nil
}
