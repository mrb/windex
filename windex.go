package windex

type Windex struct {
	logfile      *LogFile
	watcher      *Watcher
	indexer      Indexer
	log_to_index chan []byte
	Exit         chan bool
}

func New(filename string) (windex *Windex, err error) {
	logfile, err := NewLogFile(filename)
	if err != nil {
		return nil, err
	}

	watcher, err := NewWatcher()
	if err != nil {
		return nil, err
	}

	exit := make(chan bool)
	log_to_index := make(chan []byte)

	windex = &Windex{
		logfile:      logfile,
		watcher:      watcher,
		Exit:         exit,
		log_to_index: log_to_index,
	}

	go windex.startwatchloop()

	return windex, nil
}

func (windex *Windex) Watch() {
	err := windex.watcher.Watch(windex.logfile.filename)
	if err != nil {
		windex.Exit <- true
	}
}

func (windex *Windex) Index() {
}

func (windex *Windex) startwatchloop() {
	for {
		select {
		case ev := <-windex.watcher.watcher.Event:
			if ev != nil && ev.IsModify() && ev.Name == windex.logfile.filename {
				windex.logfile.moveAndFlush()
			}
		case err := <-windex.watcher.watcher.Error:
			if err != nil {
				windex.Exit <- true
			}
		}
	}

}
