package internal

import (
	"fmt"
	"io"
)

type Writer struct {
	Out io.Writer
	err error
}

func (writer *Writer) Bytes(in ...byte) *Writer {
	if writer.err != nil {
		return writer
	}

	_, err := writer.Out.Write(in)
	writer.err = err
	return writer
}

func writePrimitives[T bool | int64 | float64 | string](writer io.Writer, in ...T) error {
	for _, value := range in {
		_, err := writer.Write(fmt.Append(nil, value))
		if err != nil {
			return err
		}
	}
	return nil
}

func (writer *Writer) Bools(in ...bool) *Writer {
	if writer.err != nil {
		return writer
	}

	writer.err = writePrimitives(writer.Out, in...)
	return writer
}

func (writer *Writer) Ints(in ...int64) *Writer {
	if writer.err != nil {
		return writer
	}

	writer.err = writePrimitives(writer.Out, in...)
	return writer
}

func (writer *Writer) Floats(in ...float64) *Writer {
	if writer.err != nil {
		return writer
	}

	writer.err = writePrimitives(writer.Out, in...)
	return writer
}

func (writer *Writer) Strings(in ...string) *Writer {
	if writer.err != nil {
		return writer
	}

	writer.err = writePrimitives(writer.Out, in...)
	return writer
}

func (writer *Writer) MapFn(in map[string]string, fn func(writer io.Writer, key, value string) error) *Writer {
	for key, value := range in {
		if writer.err != nil {
			return writer
		}

		writer.err = fn(writer.Out, key, value)
	}
	return writer
}

func (writer *Writer) Map(in map[string]string, fn func(key, value string) []byte) *Writer {
	return writer.MapFn(in, func(writer io.Writer, key, value string) error {
		_, err := writer.Write(fn(key, value))
		return err
	})
}

func (writer *Writer) ArrayFn(in []string, sep []byte, fn func(writer io.Writer, i int, value string) error) *Writer {
	last := len(in) - 1
	for i, value := range in {
		if writer.err != nil {
			return writer
		}

		writer.err = fn(writer.Out, i, value)
		if i != last {
			writer.Out.Write(sep)
		}
	}
	return writer
}

func (writer *Writer) Array(in []string, sep []byte, fn func(i int, value string) []byte) *Writer {
	return writer.ArrayFn(in, sep, func(writer io.Writer, i int, value string) error {
		_, err := writer.Write(fn(i, value))
		return err
	})
}

func (writer *Writer) Complete() error {
	return writer.err
}
