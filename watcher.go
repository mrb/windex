package windex

import (
	"github.com/howeyc/fsnotify"
)

type Watcher struct {
	Watcher *fsnotify.Watcher
}

func NewWatcher() (watcher *Watcher, err error) {
	fswatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return &Watcher{
		Watcher: fswatcher,
	}, nil
}

func (watcher *Watcher) Watch(filename string) {
	watcher.Watcher.Watch(filename)
}
