package logindex

import (
	"errors"
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
}

type Index struct {
}

type ModPair struct {
	Last int64
	This int64
}

func (m *ModPair) delta() (delta int64, err error) {
	if m.This <= 0 || m.Last <= 0 {
		err = ErrIndexZero
		return 0, err
	}
	return (m.This - m.Last), nil
}

func New(filename string) (log *Log, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	info, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}

	filesize := info.Size()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if ev != nil && ev.IsModify() && ev.Name == filename {
				}
			case err := <-watcher.Error:
				if err != nil {
				}
			}
		}
	}()

	log = &Log{
		File:     file,
		FileName: filename,
		FileSize: filesize,
		Watcher:  watcher,
	}

	return log, nil
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

/*

func readAndFlush(watchfile string, pair *ModPair) {
	if pair.Last == 0 {
		pair.Last = size
	} else {
		pair.This = size

		delta := pair.delta()

		if delta > 0 {
			data := make([]byte, (delta))

			off, err := file.Seek((-1 * delta), 2)
			if err != nil {
				log.Print("Seekerr ", err)
				return
			}

			if off != 0 {
				bytesRead, err := file.Read(data)

				if err != nil {
					log.Print(err)
					return
				}
				log.Print(bytesRead, " bytes, data: ", string(data))
			}
		} else {
			return
		}
		pair.Last = size
	}
}
*/
