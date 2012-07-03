package windex

import(
	"errors"
)

var (
	ErrIndexZero       = errors.New("index error")
	ErrNoFileName      = errors.New("missing file name")
	ErrNoFile          = errors.New("missing file")
	ErrInvalidFileSize = errors.New("file size invalid")
)
