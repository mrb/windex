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

	// Use the windex generator to get a windex instance
	windex, err := windex.New(string(fname))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("*** Opening and watching: ", fname, " ***")

	// Launch the goroutine to watch our file
	go windex.Watch()

	// Launch the goroutine to index our data
	go windex.Index()

	// Call select on our exit channel. This will let us keep
	// everything moving and in the future may be of more use.
	select {
	case exit := <-windex.Exit:
		if exit {
			return
		}
	}
}
