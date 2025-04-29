package digger

import (
	"fmt"
	"os"
	"path/filepath"
)

// Config is digger config
type Config struct {
	SourceDirectory string `yaml:"sourceDirectory"`
}

// Digger general package struct
type Digger struct {
	config         Config
	processedFiles []string
}

// New constructor from Digger struct
func New(conf Config, processedFiles []string) (*Digger, error) {

	_, err := os.Stat(conf.SourceDirectory)
	if err != nil {
		return nil, err
	}

	d := Digger{
		config:         conf,
		processedFiles: processedFiles,
	}

	return &d, nil
}

// FindNewOrderNumbers get new order numbers from files
func (d *Digger) FindNewOrderNumbers() ([]string, error) {

	var newOrderNumbers []string

	entries, err := os.ReadDir(d.config.SourceDirectory)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		on, err := extractOrderNumber(filepath.Join(d.config.SourceDirectory, entry.Name()))
		if err != nil {
			return nil, err
		}
		newOrderNumbers = append(newOrderNumbers, on)
		if len(newOrderNumbers)%100 == 0 {
			fmt.Println(len(newOrderNumbers))
		}
	}

	return newOrderNumbers, nil
}
