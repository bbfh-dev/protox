package internal

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type FloatDelim struct {
	Ref   *float64
	Delim byte
}

func (step *FloatDelim) Write(writer io.Writer) error {
	_, err := fmt.Fprintf(writer, "%f%s", *step.Ref, string(step.Delim))
	return err
}

func (step *FloatDelim) Read(reader *bufio.Reader) error {
	str, err := reader.ReadString(step.Delim)
	if err != nil {
		return err
	}
	*step.Ref, err = strconv.ParseFloat(str, 64)
	return err
}
