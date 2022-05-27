package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (reader MyReader) Read(bytes []byte) (int, error) {
	i := 0
	for i < len(bytes) {
		bytes[i] = 'A'
		i++
	}
	return i, nil
}

func main() {
	reader.Validate(MyReader{})
}
