package digger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
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

	wg := sync.WaitGroup{}

	for i, entry := range entries {

		wg.Add(1)
		go func(idx int, e os.DirEntry) {
			defer wg.Done()
			orderNumber, err := extractOrderNumber(filepath.Join(d.config.SourceDirectory, e.Name()))
			if err != nil {
				panic(err)
			}
			newOrderNumbers[idx] = orderNumber
		}(i, entry)

		if i%100 == 0 {
			fmt.Println(i)
		}
	}

	wg.Wait()

	return newOrderNumbers, nil
}
