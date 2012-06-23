package windex

import (
	"errors"
	"fmt"
	"github.com/howeyc/fsnotify"
	"os"
)

var (
	ErrIndexZero       = errors.New("index error")
	ErrNoFileName      = errors.New("missing file name")
	ErrNoFile          = errors.New("missing file")
	ErrInvalidFileSize = errors.New("file size invalid")
)

type Log struct {
	FileName string
	File     *os.File
	FileSize int64
	Watcher  *fsnotify.Watcher
	Pair     *ModPair
}

type Indexer struct {
}

type ModPair struct {
	Last  int64
	This  int64
	Delta int64
}

func (m *ModPair) setDelta() (err error) {
	if m.This <= 0 || m.Last <= 0 {
		err = ErrIndexZero
		return err
	}

	m.Delta = (m.This - m.Last)

	return nil
}

func New(filename string) (log *Log, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	log = &Log{
		File:     file,
		FileName: filename,
		FileSize: 0,
		Watcher:  watcher,
		Pair:     &ModPair{0, 0, 0},
	}

	if err = log.updateFileSize(); err != nil {
		return nil, err
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

	return log, nil
}

func (log *Log) moveAndFlush() {
	if ok := log.movePair(); ok {
		log.flush()
	}
}

func (log *Log) movePair() (ok bool) {
	log.updateFileSize()

	if log.Pair.Last == 0 {
		log.Pair.Last = log.FileSize
		ok = false
	} else {
		log.Pair.This = log.FileSize
		ok = true
	}

	log.Pair.setDelta()

	log.Pair.Last = log.FileSize

	return
}

func (log *Log) updateFileSize() (err error) {
	info, err := os.Stat(log.FileName)
	if err != nil {
		return err
	}

	log.FileSize = info.Size()
	return nil
}

func (log *Log) flush() {
	delta := log.Pair.Delta
	file := log.File

	if delta > 0 {
		data := make([]byte, (delta))

		off, err := file.Seek((-1 * delta), 2)
		if err != nil {
			return
		}

		if off != 0 {
			bytesRead, err := file.Read(data)

			if err != nil {
				return
			}

			if bytesRead != 0 {
				fmt.Println(string(data))
			}
		}
	} else {
		return
	}

}

func (log *Log) watchable() (err error) {
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

func (log *Log) Watch() (err error) {
	err = log.watchable()
	if err != nil {
		return
	}

	log.Watcher.Watch(log.FileName)

	return
}
