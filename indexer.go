package windex

type Indexer interface {
  Parse() error
  Flush() error
}

type StdoutIndexer struct {

}

func (i *Indexer) Parse() (err error) {

}

func (i *Indexer) Flush() (err error) {

}
