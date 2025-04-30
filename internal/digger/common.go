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

	entries, err := os.ReadDir(d.config.SourceDirectory)
	if err != nil {
		return nil, err
	}

	newOrderNumbers := make([]string, len(entries))

	for i, entry := range entries {
		orderNumber, err := extractOrderNumber(filepath.Join(d.config.SourceDirectory, entry.Name()))
		if err != nil {
			return nil, err
		}
		newOrderNumbers[i] = orderNumber
		if i%100 == 0 {
			fmt.Println(i)
		}
	}

	return newOrderNumbers, nil
}
