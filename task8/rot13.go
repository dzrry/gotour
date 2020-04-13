package task8

import (
	"io"
)

type Rot13Reader struct {
	R io.Reader
}

func (r Rot13Reader) Read(p []byte) (n int, err error) {

	n, err = r.R.Read(p)
	for i := range p {
		switch {
		case p[i] >= 'A' && p[i] < 'N':
			fallthrough
		case p[i] >= 'a' && p[i] < 'n':
			p[i] += 13
		case p[i] > 'M' && p[i] <= 'Z':
			fallthrough
		case p[i] > 'm' && p[i] <= 'z':
			p[i] -= 13
		}
	}
	return
}
