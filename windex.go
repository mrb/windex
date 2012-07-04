package windex

import (
	"time"
)

type Windex struct {
	logfile *LogFile
	watcher *Watcher
	indexer Indexer
	LogData chan []byte
	Exit    chan bool
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

	indexer := NewStdoutIndexer()

	exit := make(chan bool)
	log_data := make(chan []byte)

	windex = &Windex{
		logfile: logfile,
		watcher: watcher,
		indexer: indexer,
		Exit:    exit,
		LogData: log_data,
	}

	go windex.startwatchloop()

	return windex, nil
}

func (windex *Windex) Watch() {
	for {
		time.Sleep(50 * time.Millisecond)
		windex.watcher.Watch(windex.logfile.Filename)
	}
}

func (windex *Windex) Index() {
	for {
		windex.indexer.Parse()
		windex.indexer.Flush(windex.LogData)
	}
}

func (windex *Windex) startwatchloop() {
	for {
		select {
		case ev := <-windex.watcher.Watcher.Event:
			if ev != nil && ev.IsModify() && ev.Name == windex.logfile.Filename {
				windex.logfile.Flush(windex.LogData)
			}
		case err := <-windex.watcher.Watcher.Error:
			if err != nil {
				windex.Exit <- true
			}
		}
	}

}
