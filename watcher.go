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
