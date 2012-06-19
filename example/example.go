package main

import (
	"github.com/mrb/logindex"
	"log"
	"time"
)

func main() {
	logg, err := logindex.New("logs.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(logg, " ", err)

	for {
		time.Sleep(500 * time.Millisecond)
		err = logg.Watch()
		if err != nil {
			log.Print(err)
		}
	}
}
