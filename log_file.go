package windex

import (
	"os"
)

type LogFile struct {
	filename string
	file     *os.File
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

	defer file.Close()

	return &LogFile{
		file:     file,
		filename: filename,
		filesize: 0,
		cursor:   &LogFileCursor{0, 0, 0},
	}, nil
}

func (m *LogFileCursor) setDelta() (err error) {
	if m.this <= 0 || m.last <= 0 {
		err = ErrIndexZero
		return err
	}

	m.delta = (m.this - m.last)

	return nil
}

func (log *LogFile) moveAndFlush() {
	if ok := log.movePair(); ok {
		log.flush()
	}
}

func (log *LogFile) movePair() (ok bool) {
	log.updateFileSize()

	if log.cursor.last == 0 {
		log.cursor.last = log.filesize
		ok = false
	} else {
		log.cursor.this = log.filesize
		ok = true
	}

	log.cursor.setDelta()

	log.cursor.last = log.filesize

	return
}

func (log *LogFile) updateFileSize() (err error) {
	info, err := os.Stat(log.filename)
	if err != nil {
		return err
	}

	log.filesize = info.Size()
	return nil
}

//
func (log *LogFile) flush() {
	delta := log.cursor.delta
	file := log.file

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
				//log.Indexer.flush()
				os.Stdout.Write(data)
			}
		}
	} else {
		return
	}

}
