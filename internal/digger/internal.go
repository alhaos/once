package digger

import (
	"bytes"
	"os"
)

const sep byte = 0x0d

// extractOrderNumber extract order number from file
func extractOrderNumber(filename string) (string, error) {

	var orderNumber []byte

	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	lines := bytes.Split(content, []byte{sep})

	for _, line := range lines {
		if bytes.HasPrefix(line, []byte("ORC")) {
			fields := bytes.Split(line, []byte("|"))

			if len(fields) < 2 {
				continue
			}

			if len(fields[2]) > 0 {
				orderNumber = fields[2]
			}
		}
	}

	return string(orderNumber), err
}
