## Windex: Watch and Index

Windex is a library for concurrently watching log files and indexing their data.

A `Watcher` provides cross-platform file system notifications via
`github.com/howeyc/fsnotify`. These events help keep track of changes to the
size of a file over time. This information is stored, and data from the `tail` of
the file is sent via a channel to an external source, known as an`Indexer`.

The default `Indexer` flushes to `stdout`, effectively making `windex` a
programatically accessible `tail -f` in its default state. However it can be
much more powerful than that.

`Indexer` is an interface that contains two methods:

```go
type Indexer interface {
	Parse(chan []byte) ([]byte, error)
	Flush([]byte) error
}
```

The `stdout` indexer does nothing with `Parse(chan []byte)`, and simply prints
with `Flush([]byte)`. Since `Parse` has access to the data, it will be a 
powerful step toward implementing `Indexers` which can parse and flush data to
any number of a variety of external sources: Redis, Postgres, Statsd, etc.

Here's an example naive implementation of a Redis Indexer:

```go
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
```

### Usage

```go
// Use the windex generator to get a windex instance
windex, err = windex.New("logfile01.log")

// Launch the goroutine to watch our file
go windex.Watch()

// Launch the goroutine to index our data. The default indexer
// flushes data to stdout.
go windex.Index()

// Call select on our exit channel. This will let us keep
// everything moving and in the future may be of more use.
select {
case exit := <-windex.Exit:
	if exit {
		return
	}
}
```

### Example

The above is paraphrased from the example program in `example/stdoutindexer/example.go`. There is an example that uses the above Redis Indexer in `example/redisindexer/example.go`

### Credits

windex is (c) Michael R. Bernstein, 2012

### License

windex is distributed under the MIT License, see `LICENSE` file for details.
