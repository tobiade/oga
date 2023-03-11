package lang

import (
	"fmt"
	"os"
)

type SourceProvider interface {
	Get() (string, error)
}

// FileSourceProvider gets oga source code from a file
type FileSourceProvider struct {
	filepath string
}

func NewDefaultSourceProvider(filepath string) *FileSourceProvider {
	return &FileSourceProvider{filepath: filepath}
}

func (p *FileSourceProvider) Get() (string, error) {
	b, err := os.ReadFile(p.filepath)
	if err != nil {
		return "", fmt.Errorf("unable to read file [%s]: %w", p.filepath, err)
	}
	return string(b), nil
}
