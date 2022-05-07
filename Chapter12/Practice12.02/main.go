package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	logFile, err := os.OpenFile(
		fmt.Sprintf("log-%v.txt", time.Now().Format("2006-01-02")),
		os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "log:", log.Ldate|log.Lmicroseconds|log.Llongfile)
	logger.Println("log message")
}
