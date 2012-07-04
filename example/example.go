package main

import (
	"github.com/mrb/windex"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Please supply a log file to watch")
	}

	fname := os.Args[1]

	windex, err := windex.New(string(fname))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Opening and watching", fname)

  go windex.Watch()
  go windex.Index()

  <-windex.Exit
}
