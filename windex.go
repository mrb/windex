package windex

import (
	"log"
)

type Windex struct {
	logfile      *LogFile
	watcher      *Watcher
	indexer      Indexer
	log_to_index chan []byte
	Exit         chan bool
}

func New(filename string) (windex *Windex, err error) {
	logfile, err := NewLogFile(filename)
	if err != nil {
		return nil, err
	}

	watcher, err := NewWatcher()
	if err != nil {
		return nil, err
	}

	exit := make(chan bool)
	log_to_index := make(chan []byte)

	windex = &Windex{
		logfile:      logfile,
		watcher:      watcher,
		Exit:         exit,
		log_to_index: log_to_index,
	}

	go windex.startwatchloop()

	return windex, nil
}

func (windex *Windex) Watch() {
	windex.watcher.Watch(windex.logfile.Filename)
}

func (windex *Windex) Index() {
}

func (windex *Windex) startwatchloop() {
	for {
		select {
		case ev := <-windex.watcher.Watcher.Event:
			if ev != nil && ev.IsModify() && ev.Name == windex.logfile.Filename {
				//windex.logfile.moveAndFlush()
				log.Print(ev)
			}
		case err := <-windex.watcher.Watcher.Error:
			if err != nil {
				windex.Exit <- true
			}
		}
	}

}
