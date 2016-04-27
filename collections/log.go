package collections

import (
	"os"
	"log"
)

func Log(err error){
	f, err_opening := os.OpenFile("/var/log/monagent.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666)
	if err_opening != nil{
		panic(err_opening)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Printf("Error: %v", err)
}