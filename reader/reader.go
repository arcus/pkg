package reader

import (
	"bytes"
	"io"
)

var bom = []byte{0xef, 0xbb, 0xbf}

// UniversalReader wraps an io.Reader to replace carriage returns with newlines.
// This is used with the csv.Reader so it can properly delimit lines.
type UniversalReader struct {
	r io.Reader
}

func (r *UniversalReader) Read(buf []byte) (int, error) {
	n, err := r.r.Read(buf)

	// Detect and remove BOM.
	if bytes.HasPrefix(buf, bom) {
		copy(buf, buf[len(bom):])
		n -= len(bom)
	}

	// Replace carriage returns with newlines
	for i, b := range buf {
		if b == '\r' {
			buf[i] = '\n'
		}
	}

	return n, err
}

func (r *UniversalReader) Close() error {
	if rc, ok := r.r.(io.Closer); ok {
		return rc.Close()
	}
	return nil
}

func NewUniversalReader(r io.Reader) *UniversalReader {
	return &UniversalReader{r}
}
