package main

import (
	"github.com/howeyc/fsnotify"
	"log"
	"os"
	"time"
)

type ModPair struct {
	Last int64
	This int64
}

func (m *ModPair) delta() (delta int64) {
	return (m.This - m.Last)
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	log.Print(watcher)

	watchfile := "logs.txt"

	pair := &ModPair{0, 0}

	file, err := os.Open(watchfile)
	if err != nil {
		log.Fatal(err, "Could not open file")
	}

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if ev != nil && ev.IsModify() && ev.Name == watchfile {
					info, _ := os.Stat(watchfile)
					size := info.Size()

					if pair.Last == 0 {
						pair.Last = size
					} else {
						pair.This = size

						delta := pair.delta()

						if delta > 0 {
							data := make([]byte, (delta))

							off, err := file.Seek((-1 * delta), 2)
							if err != nil {
								log.Print("Seekerr ", err)
								return
							}

							if off != 0 {
								bytesRead, err := file.Read(data)

								if err != nil {
									log.Print(err)
									return
								}
								log.Print(bytesRead, " bytes, data: ", string(data))
							}
						} else {
							return
						}

						pair.Last = size
					}

				}
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	for {
		time.Sleep(500 * time.Millisecond)
		err = watcher.Watch("logs.txt")
		if err != nil {
			log.Fatal(err)
		}
	}
}
