package windex

import (
	"github.com/howeyc/fsnotify"
)

type Watcher struct {
	watcher *fsnotify.Watcher
}

func NewWatcher() (watcher *Watcher, err error) {
	return &Watcher{
		watcher: &fsnotify.Watcher{},
	}, nil
}

func (watcher *Watcher) Watch(filename string) (err error) {
	watcher.Watch(filename)

	return
}

/*
go func() {
  for {
	select {
	case ev := <-watcher.Event:
		if ev != nil && ev.IsModify() && ev.Name == filename {
			log.moveAndFlush()
		}
	case err := <-watcher.Error:
		if err != nil {
		}
	}
  }
}()
*/
/*

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	logfile := &LogFile{}

	if err = log.updateFileSize(); err != nil {
		return nil, err
	}

*/
