package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/mrb/windex"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Please supply at least one log file to watch")
	}

	windexes := []*windex.Windex{}

	redis_indexer := NewRedisIndexer("localhost:6379")

	for i, name := range os.Args {
		if i > 0 {
			thiswindex, _ := windex.New(name, redis_indexer)
			windexes = append(windexes, thiswindex)
			log.Print(name)
		}
	}

	for _, thiswindex := range windexes {
		log.Println("*** Opening and watching: ", thiswindex.Filename(), " ***")
		go thiswindex.Watch()
		go thiswindex.Index()
	}

	select {
	case exit := <-windexes[0].Exit:
		if exit {
			return
		}
	}
}

type RedisIndexer struct {
	address string
}

func (i *RedisIndexer) Parse(log_data chan []byte) (parsed_log_data []byte, err error) {
	parsed_log_data = <-log_data
	return parsed_log_data, nil
}

func (i *RedisIndexer) Flush(parsed_log_data []byte) (err error) {
	redis, err := redis.Dial("tcp", i.address)

	if err != nil {
		log.Fatal(err)
	}

	redis.Do("lpush", "errorz", parsed_log_data)

	log.Print("Data Pushed: ", len(parsed_log_data), " bytes")

	return
}

func NewRedisIndexer(address string) (stdout *RedisIndexer) {
	return &RedisIndexer{
		address: address,
	}
}
