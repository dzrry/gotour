package task7

import "io"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, io.ErrShortBuffer
	}
	for idx := range b {
		b[idx] = 'A'
	}
	return len(b), nil
}