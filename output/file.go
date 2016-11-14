package output

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// FileWriter writes the output to file
type FileWriter struct {
	targetFile string
	fileMode   os.FileMode
}

// NewFileWriter creates new FileWriter what writes to targetFile
func NewFileWriter(targetFile string) *FileWriter {
	return &FileWriter{
		targetFile: targetFile,
		fileMode:   0600,
	}
}

func (w *FileWriter) write(output string) error {
	targetDir := filepath.Dir(w.targetFile)
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		return fmt.Errorf("Directory %s of target file does not exist", targetDir)
	}
	return ioutil.WriteFile(w.targetFile, []byte(output), w.fileMode)
}
