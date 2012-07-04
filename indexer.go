package windex

import (
	logger "log"
)

type Indexer interface {
	Parse() error
	Flush(chan []byte) error
}

type StdoutIndexer struct {
}

func (i *StdoutIndexer) Parse() (err error) {
	return
}

func (i *StdoutIndexer) Flush(log_data chan []byte) (err error) {
	logger.Print(<-log_data)
	return
}

func NewStdoutIndexer() (stdout *StdoutIndexer) {
	return &StdoutIndexer{}
}
