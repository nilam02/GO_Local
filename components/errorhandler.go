package components

import "log"

// HandleError logs the error and returns it as an error type
func HandleError(err error, message string) error {
	if err != nil {
		log.Println(message, err)
		return err
	}
	return nil
}
