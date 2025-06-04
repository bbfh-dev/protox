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
	str := strconv.FormatFloat(*step.Ref, 'f', -1, 64)
	_, err := fmt.Fprintf(writer, "%s%s", str, string(step.Delim))
	return err
}

func (step *FloatDelim) Read(reader *bufio.Reader) error {
	str, err := reader.ReadString(step.Delim)
	if err != nil {
		return err
	}
	*step.Ref, err = strconv.ParseFloat(str[:len(str)-1], 64)
	return err
}
