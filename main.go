package main

import (
	"log"
	"os"
)

var Log *log.Logger

func logger() {
	logFile, e := os.OpenFile("C:\\Users\\SJH\\Desktop\\log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if e != nil {
		log.Panic("Opening Log File is Error!", e)
	}
	Log = log.New(logFile, "Log ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	logger()
	Log.Println("start server")
	Home()
}
