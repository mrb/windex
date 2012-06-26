package windex

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"testing"
	"time"

	"github.com/bmizerany/assert"
	"github.com/mrb/windex"
)

var (
	LogLine         = []byte("[2012-06-23T00:00:00+00:00] Boom  Boom  Boom  Boom  Boom  Boom  Boom  Boom  Boom\n")
	Lines           = 20
	TestLogFileName = string("windex_test_log.log")
	InputBuffer     []byte
	OutputBuffer    []byte
	WindexLog       *windex.Log
)

func setupInputBuffer(t *testing.T) {
	InputBuffer = make([]byte, 0)
	assert.T(t, InputBuffer != nil)

	for i := 0; i < Lines; i++ {
		InputBuffer = append(InputBuffer, LogLine...)
	}
	assert.T(t, len(InputBuffer) == (Lines*len(LogLine)))

	return
}

func setupOutputBuffer(t *testing.T) {
	OutputBuffer = make([]byte, 0)
	assert.T(t, OutputBuffer != nil)
	return
}

func createTestLogFile(t *testing.T) {
	file, cerr := os.OpenFile(TestLogFileName, os.O_CREATE, 0777)
	defer file.Close()
	assert.T(t, cerr == nil)
	assert.T(t, file != nil)
}

func writeInputToLogFile(t *testing.T) {
	sep := []byte("\n")
	lines := bytes.Split(InputBuffer, sep)

	assert.T(t, lines != nil)
	assert.T(t, len(lines) == Lines+1)

	var file, err = os.OpenFile(TestLogFileName, os.O_RDWR, 0777)
	assert.T(t, err == nil)
	assert.T(t, file != nil)

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		line = append(line, byte('\n'))
		b, err := file.Write([]byte(line))
		assert.T(t, err == nil)
		assert.T(t, b >= 0)
	}

	info, err := os.Stat(TestLogFileName)
	assert.T(t, info != nil)
	assert.T(t, err == nil)
	assert.T(t, info.Size() == int64(Lines*len(LogLine)+1))
}

func deleteLogFile(t *testing.T) {
	err := os.Remove(TestLogFileName)
	assert.T(t, err == nil)
}

func TestOutputLength(t *testing.T) {
	setupInputBuffer(t)
	setupOutputBuffer(t)
	createTestLogFile(t)
	defer deleteLogFile(t)

	WindexLog, err := windex.New(TestLogFileName)
	assert.T(t, err == nil)

	OutputBuffer = make([]byte, int64(5*len(LogLine)))

	r := bufio.NewReader(os.Stdout)

	log.Print(WindexLog)

	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			log.Print(WindexLog.FileSize)
			err = WindexLog.Watch()
			if err != nil {
				log.Print(err)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		writeInputToLogFile(t)
	}

	b, err := r.Read(OutputBuffer)
	assert.T(t, err == nil)
	assert.T(t, b > 0)

	log.Print(OutputBuffer)
}
