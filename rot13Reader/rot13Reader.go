package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(bytes []byte) (int, error) {
	reader.r.Read(bytes)
	i := 0
	for i < len(bytes) {
		switch {
		case (bytes[i] >= 'A' && bytes[i] <= 'M') || (bytes[i] >= 'a' && bytes[i] <= 'm'):
			{
				bytes[i] += 13
			}
		case (bytes[i] >= 'N' && bytes[i] <= 'Z') || (bytes[i] >= 'n' && bytes[i] <= 'z'):
			{
				bytes[i] -= 13
			}
		}
		i++
	}
	return i, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
