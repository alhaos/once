package digger

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

const sep byte = 0x0d

func extractOrderNumber(path string) (string, error) {

	var file, err = os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {

		b, err := reader.ReadBytes(0x0d)
		if err != nil {
			return "", err
		}

		if !bytes.HasPrefix(b, []byte("ORC")) {
			continue
		}

		var fields = bytes.Split(b, []byte("|"))

		if len(fields) < 2 {
			continue
		}

		if len(fields[2]) > 0 {
			return string(fields[2]), err
		}
	}
	return "", err
}
