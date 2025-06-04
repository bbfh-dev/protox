package internal

import (
	"bufio"
	"io"
)

type Byte struct {
	Ref *byte
}

func (step *Byte) Write(writer io.Writer) (err error) {
	_, err = writer.Write([]byte{*step.Ref})
	return err
}

func (step *Byte) Read(reader *bufio.Reader) (err error) {
	*step.Ref, err = reader.ReadByte()
	return err
}

type ByteDelim struct {
	Value *[]byte
	Delim byte
}

func (step *ByteDelim) Write(writer io.Writer) (err error) {
	_, err = writer.Write(*step.Value)
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte{step.Delim})
	return err
}

func (step *ByteDelim) Read(reader *bufio.Reader) error {
	slice, err := reader.ReadBytes(step.Delim)
	*step.Value = slice[:len(slice)-1]
	return err
}
