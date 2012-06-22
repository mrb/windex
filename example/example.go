package main

import (
	"github.com/mrb/windex"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Please supply a log file to watch")
	}

	fname := os.Args[1]

	indexed_log, err := windex.New(string(fname))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Opening and watching", fname)

	for {
		time.Sleep(500 * time.Millisecond)
		err = indexed_log.Watch()
		if err != nil {
			log.Print(err)
		}
	}
}
