package windex

type Indexer interface {
	Parse() error
	Flush() error
}

type StdoutIndexer struct {
}

func (i *StdoutIndexer) Parse() (err error) {
	return
}

func (i *StdoutIndexer) Flush() (err error) {
	return
}
