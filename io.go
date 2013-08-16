package mtf

import "io"

type writer struct {
	w io.Writer
	a []byte
}

// Create a proxying writer that performs MTF on data given to it. The data is
// modified in-place when writing.
func NewWriter(w io.Writer, alphabet []byte) io.Writer {
	if alphabet == nil {
		alphabet = DefaultAlphabet()
	}
	return writer{w, alphabet}
}

func (w writer) Write(b []byte) (n int, err error) {
	MTF(b, w.a)
	return w.w.Write(b)
}

type reader struct {
	r io.Reader
	a []byte
}

// Create a new proxying reader that performs UnMTF on data read.
func NewReader(r io.Reader, alphabet []byte) io.Reader {
	if alphabet == nil {
		alphabet = DefaultAlphabet()
	}
	return reader{r, alphabet}
}

func (r reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)
	UnMTF(b[:n], r.a)
	return n, err
}
