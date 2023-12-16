package utils

import (
	"log"
)

func Log(src string, message string, error error) {
	if error != nil {
		log.Println("[ERROR] [" + src + "] " + message + ": " + error.Error())
	} else {
		log.Println("[" + src + "] " + message)
	}
}
