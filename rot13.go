package main

import (
	"io"
	"os"
	"strings"
)

const lowerA byte = 97
const lowerZ byte = 122

const upperA byte = 65
const upperZ byte = 90

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(p []byte) (n int, err error) {
	bytesWritten, err := reader.r.Read(p)

	if err != nil {
		return bytesWritten, err
	}

	for i := 0; i < len(p); i++ {
		var b byte = p[i]
		var min byte

		if b <= upperZ {
			min = upperA
		} else {
			min = lowerA
		}

		valueAdd := (b - min + 13) % 26

		p[i] = min + valueAdd
	}

	return bytesWritten, nil
}

func main() {
	s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
