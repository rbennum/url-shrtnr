package utils

import (
	"io"
	"log"
	"os"
)

func InitLogger() error {
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		os.Mkdir("log", os.ModePerm)
	}
	logFile, err := os.OpenFile("log/log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	writer := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(writer)
	return err
}
