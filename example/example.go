package main

import (
	"github.com/mrb/windex"
	"log"
	"time"
)

func main() {
	indexed_log, err := windex.New("logs.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(indexed_log, " ", err)

	for {
		time.Sleep(500 * time.Millisecond)
		err = indexed_log.Watch()
		if err != nil {
			log.Print(err)
		}
	}
}
