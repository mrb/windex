package windex

import (
	"fmt"
)

type Indexer interface {
	Parse(chan []byte) ([]byte, error)
	Flush([]byte) error
}

type StdoutIndexer struct {
}

func (i *StdoutIndexer) Parse(log_data chan []byte) (parsed_log_data []byte, err error) {
	parsed_log_data = <-log_data
	return parsed_log_data, nil
}

func (i *StdoutIndexer) Flush(parsed_log_data []byte) (err error) {
	fmt.Printf("%s", parsed_log_data)
	return
}

func NewStdoutIndexer() (stdout *StdoutIndexer) {
	return &StdoutIndexer{}
}
