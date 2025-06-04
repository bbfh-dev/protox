package internal

import (
	"bufio"
	"io"
)

type String struct {
	Ref  *string
	Size int
}

func (step *String) Write(writer io.Writer) (err error) {
	_, err = writer.Write([]byte(*step.Ref))
	return err
}

func (step *String) Read(reader *bufio.Reader) (err error) {
	slice := make([]byte, step.Size)
	_, err = reader.Read(slice)
	if err != nil {
		return err
	}

	*step.Ref = string(slice)
	return
}

type StringDelim struct {
	Ref   *string
	Delim byte
}

func (step *StringDelim) Write(writer io.Writer) (err error) {
	_, err = writer.Write([]byte(*step.Ref))
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte{step.Delim})
	return err
}

func (step *StringDelim) Read(reader *bufio.Reader) error {
	str, err := reader.ReadString(step.Delim)
	*step.Ref = str[:len(str)-1]
	return err
}
