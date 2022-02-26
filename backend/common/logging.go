package common

import (
	"fmt"
	"log"
)

func LogFatal(tag string, message string) {
	log.Fatalln(fmt.Sprintf("[%s] : %s", tag, message))
}

func LogPrintln(tag string, message string) {
	log.Println(fmt.Sprintf("[%s] : %s", tag, message))
}
