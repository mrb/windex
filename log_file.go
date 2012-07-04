package windex

import (
	"os"
)

type LogFile struct {
	Filename string
	File     *os.File
	filesize int64
	cursor   *LogFileCursor
}

type LogFileCursor struct {
	last  int64
	this  int64
	delta int64
}

func NewLogFile(filename string) (log_file *LogFile, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	log_file = &LogFile{
		File:     file,
		Filename: filename,
		filesize: 0,
		cursor:   &LogFileCursor{0, 0, 0},
	}

	if err = log_file.updateFileSize(); err != nil {
		return nil, err
	}

	return log_file, nil
}

func (cursor *LogFileCursor) setDelta() (err error) {
	if cursor.this <= 0 || cursor.last <= 0 {
		err = ErrIndexZero
		return err
	}

	cursor.delta = (cursor.this - cursor.last)

	return nil
}

func (log_file *LogFile) Flush(log_data chan []byte) {
	if ok := log_file.movePair(); ok {
		log_file.flush(log_data)
	}
}

func (log_file *LogFile) movePair() (ok bool) {
	log_file.updateFileSize()

	if log_file.cursor.last == 0 {
		log_file.cursor.last = log_file.filesize
		ok = false
	} else {
		log_file.cursor.this = log_file.filesize
		ok = true
	}

	log_file.cursor.setDelta()

	log_file.cursor.last = log_file.filesize

	return
}

func (log_file *LogFile) updateFileSize() (err error) {
	info, err := os.Stat(log_file.Filename)
	if err != nil {
		return err
	}

	log_file.filesize = info.Size()
	return nil
}

// TODO: Maybe break up, definitely handle errors better
func (log_file *LogFile) flush(log_data chan []byte) {
	delta := log_file.cursor.delta
	file := log_file.File

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
				log_data <- data
			}
		}
	} else {
		return
	}
}
