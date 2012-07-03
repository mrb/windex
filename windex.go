package windex

import (
	"os"
)

type Windex struct {
  watcher *Watcher
  logfile *LogFile
  indexer *Indexer
  logchan chan []byte
}

/*

watched_index, err = windex.New("logfile01.log")
watched_index.Watch()
watched_index.Index()
// or .Index(StdoutIndex) where StdoutIndex implements
// Index interface

Windex methods orchestrate between logfile and indexer,
getting signals from watcher to know when to act

[]byte channel between logfile and indexer

*/
func New(filename string) (windex *Windex, err error) {
	Watcher     *fsnotify.Watcher
	
        if err != nil {
		return nil, err
	}

	defer file.Close()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	logfile := &LogFile{}

	if err = log.updateFileSize(); err != nil {
		return nil, err
	}

	return log, nil
}


