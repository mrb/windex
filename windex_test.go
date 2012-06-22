package windex

import (
	//"github.com/mrb/windex"
	"github.com/bmizerany/assert"
	"log"
	"os"
	"testing"
)

func setupInputFile(t *testing.T) (file *os.File) {
	file = os.NewFile(uintptr(os.O_CREATE), "input.test")
	assert.T(t, file != nil)
	return
}

func setupOutputFile(t *testing.T) (file *os.File) {
	file = os.NewFile(uintptr(os.O_CREATE), "output.test")
	assert.T(t, file != nil)
	return
}

func TestFileSetup(t *testing.T) {
	input := setupInputFile(t)
	assert.T(t, input != nil)

	output := setupOutputFile(t)
	assert.T(t, output != nil)
}

func TestWriteToInputFile(t *testing.T) {
	input := setupInputFile(t)

	for i := 0; i < 10; i++ {
		input.WriteString("Test\n")
	}

	log.Print(os.Stat("intput.test"))
}
