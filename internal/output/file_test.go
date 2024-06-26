package output

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	file, createErr := os.CreateTemp("", "example")
	assert.NoError(t, createErr, "Unable to create temp file")

	writer := NewFileWriter(file.Name(), 0o640)

	_ = writer.write("foobar")
	_ = writer.write("foobar-second-time")

	fileBytes, readErr := os.ReadFile(file.Name())
	assert.NoError(t, readErr, "Unable to read temp file")

	info, err := os.Stat(file.Name())
	if err != nil || info == nil {
		assert.Fail(t, "Unable to get file info")

		return
	}

	assert.Equal(t, 0o640, int(info.Mode().Perm()), "File permission wasn't set as expected")
	assert.Equal(t, "foobar-second-time", string(fileBytes), "FileWriter didnt wrote expected output to file")
}

func TestFailIfNoDirectory(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	writer := NewFileWriter("do/not/exist/my-file.txt", 0o640)

	err := writer.write("foobar")
	assert.Error(t, err, "Did not return error if directory doesn't exist")
}
