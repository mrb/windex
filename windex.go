package windex

type Windex struct {
	logfile      *LogFile
	watcher      *Watcher
	indexer      *Indexer
	log_to_index chan []byte
	exit         chan bool
}

/*

windex, err = windex.New("logfile01.log")
err = windex.Watch()
err = windex.Index()
// or .Index(StdoutIndex) where StdoutIndex implements
// Index interface

exit <- windex.exit

Windex methods orchestrate between logfile and indexer,
getting signals from watcher to know when to act

[]byte channel between logfile and indexer
bool channel between windex and the outside world

*/
func New(filename string) (windex *Windex, err error) {
	logfile, err := NewLogFile(filename)
	if err != nil {
		return nil, err
	}

	watcher, err := NewWatcher()
	if err != nil {
		return nil, err
	}

	return &Windex{
		logfile: logfile,
		watcher: watcher,
	}, nil
}

func (windex *Windex) Watch() (err error) {

}

func (windex *Windex) Index() (err error) {

}
