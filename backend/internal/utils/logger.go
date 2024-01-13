package utils

import (
	"log"
)

// Log Formats and logs the message to the console.
// src is the source of the log message.
// message is the message to be logged.
// error is the error to be logged, if any. If not nil, message will be prepended with "[ERROR] ".
func Log(src string, message string, error error) {
	if error != nil {
		log.Println("[ERROR] [" + src + "] " + message + ": " + error.Error())
	} else {
		log.Println("[" + src + "] " + message)
	}
}
