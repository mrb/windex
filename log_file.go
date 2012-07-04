package windex

import (
	"os"
)

type LogFile struct {
	FileName string
	File     *os.File
	FileSize int64
	Cursor   *LogFileCursor
}

type LogFileCursor struct {
	Last  int64
	This  int64
	Delta int64
}

func NewLogFile(filename string) (log_file *LogFile, err error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	return &LogFile{
		File:     file,
		FileName: filename,
		FileSize: 0,
		Pair:     &LogFileCursor{0, 0, 0},
	}
}

func (m *LogFileCursor) setDelta() (err error) {
	if m.This <= 0 || m.Last <= 0 {
		err = ErrIndexZero
		return err
	}

	m.Delta = (m.This - m.Last)

	return nil
}

func (log *LogFile) moveAndFlush() {
	if ok := log.movePair(); ok {
		log.flush()
	}
}

func (log *LogFile) movePair() (ok bool) {
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

func (log *LogFile) updateFileSize() (err error) {
	info, err := os.Stat(log.FileName)
	if err != nil {
		return err
	}

	log.FileSize = info.Size()
	return nil
}

//
func (log *LogFile) flush() {
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
				log.Indexer.flush()
				os.Stdout.Write(data)
			}
		}
	} else {
		return
	}

}
