package components

import "log"

func LogMessage(logger *log.Logger, message string) {
	logger.Println(message)
}