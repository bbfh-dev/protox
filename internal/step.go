package internal

import (
	"bufio"
	"io"
)

type Step interface {
	Write(io.Writer) error
	Read(*bufio.Reader) error
}
