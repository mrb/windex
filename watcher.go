package windex

import(
	"github.com/howeyc/fsnotify"
)

type Watcher struct {

}

func (log *LogFile) watchable() (err error) {
	if log.File == nil {
		err = ErrNoFile
	}

	if log.FileName == "" {
		err = ErrNoFileName
	}

	if log.FileSize < 0 {
		err = ErrInvalidFileSize
	}

	return
}

func (log *LogFile) Watch() (err error) {
	err = log.watchable()
	if err != nil {
		return
	}

	log.Watcher.Watch(log.FileName)

	return
}

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


