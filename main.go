package main

import (
	"log"
)

func main() {
	ConfigLoad()
	// Бесконечный цикл, чтобы программа продолжала работу и не завершалась сразу
	webRunServer()
	select {}
}

func checkFatalError(err error) {
	if err == nil {
		return
	}
	log.Print("Fatal Error")
	log.Fatalln(err.Error())
}
