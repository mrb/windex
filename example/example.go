package main

import (
	"github.com/mrb/windex"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Please supply at least one log file to watch")
	}

	windexes := []*windex.Windex{}

	for i, name := range os.Args {
		if i > 0 {
			thiswindex, _ := windex.New(name)
			windexes = append(windexes, thiswindex)
			log.Print(name)
		}
	}

	for _, thiswindex := range windexes {
		log.Println("*** Opening and watching: ", thiswindex.Filename(), " ***")
		go thiswindex.Watch()
		go thiswindex.Index()
	}

	select {
	case exit := <-windexes[0].Exit:
		if exit {
			return
		}
	}
}
