package main

import (
	"io"
	"os"
	"strings"
)

const lowercaseA byte = 97
const lowercaseZ byte = 122

const uppercaseA byte = 65
const uppercaseZ byte = 90

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(p []byte) (n int, err error) {
	bytesToWrite, err := reader.r.Read(p)

	if err != nil {
		return bytesToWrite, err
	}

	for i := 0; i < len(p); i++ {
		var b byte = p[i]
		var min byte

		if b <= uppercaseZ {
			min = uppercaseA
		} else {
			min = lowercaseA
		}

		valueAdd := (b - min + 13) % 26

		p[i] = min + valueAdd
	}

	return bytesToWrite, nil
}

func main() {
	s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
