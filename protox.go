package protox

import (
	"io"

	internal "github.com/bbfh-dev/protox/internal"
)

func NewReader(reader io.Reader) *internal.Reader {
	return &internal.Reader{
		In: reader,
	}
}

func NewWriter(writer io.Writer) *internal.Writer {
	return &internal.Writer{
		Out: writer,
	}
}
