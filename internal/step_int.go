package internal

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"strconv"
)

type Int struct {
	Ref *int64
}

func (step *Int) Write(writer io.Writer) error {
	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, uint64(*step.Ref))
	return nil
}

func (step *Int) Read(reader *bufio.Reader) (err error) {
	data := make([]byte, 8)
	_, err = reader.Read(data)
	if err != nil {
		return err
	}
	*step.Ref = int64(binary.LittleEndian.Uint64(data))
	return nil
}

type IntDelim struct {
	Ref   *int64
	Delim byte
}

func (step *IntDelim) Write(writer io.Writer) error {
	_, err := fmt.Fprintf(writer, "%d%s", *step.Ref, string(step.Delim))
	return err
}

func (step *IntDelim) Read(reader *bufio.Reader) (err error) {
	str, err := reader.ReadString(step.Delim)
	if err != nil {
		return err
	}
	*step.Ref, err = strconv.ParseInt(str, 10, 64)
	return err
}
