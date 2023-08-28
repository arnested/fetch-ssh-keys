package output

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestStdout(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	writer := &StdoutWriter{}
	err := writer.write("whoop")

	assert.NoError(t, err, "Unable to write to stdout")
}
