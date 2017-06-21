package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rotReader *rot13Reader) Read(b []byte) (int, error) {
	n, err := rotReader.r.Read(b)

	for i := 0; i < n; i++ {
		if (b[i] >= 'A' && b[i] < 'N') || (b[i] >= 'a' && b[i] < 'n') {
			b[i] += 13
		} else if (b[i] > 'M' && b[i] <= 'Z') || (b[i] > 'm' && b[i] <= 'z') {
			b[i] -= 13
		}
	}

	return n, err
}

func main {
	s := strings.NewReader("lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.stdout, &r)
}