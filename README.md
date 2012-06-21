## Windex: Watch and Index

```go
indexed_log, err := windex.New("log.txt")

for {
		time.Sleep(500 * time.Millisecond)
		err = indexed_log.Watch()
		if err != nil {
			log.Print(err)
		}
	}
```

### Coming Soon

* Indexing
