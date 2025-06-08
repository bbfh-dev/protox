package internal

import (
	"bufio"
	"fmt"
	"io"
)

type StringArray struct {
	Value *[]string
	Comma byte
	Delim byte
}

func (step *StringArray) Write(writer io.Writer) error {
	for _, value := range *step.Value {
		_, err := fmt.Fprintf(writer, "%s%s", value, string(step.Comma))
		if err != nil {
			return err
		}
	}
	_, err := writer.Write([]byte{step.Delim})
	return err
}

func (step *StringArray) Read(reader *bufio.Reader) error {
	for {
		peek, err := reader.Peek(1)
		if err == nil && peek[0] == step.Delim {
			reader.Discard(1)
			break
		}

		value, err := reader.ReadString(step.Comma)
		if err != nil {
			return err
		}

		*step.Value = append(*step.Value, value[:len(value)-1])
	}

	return nil
}
