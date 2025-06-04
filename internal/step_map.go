package internal

import (
	"bufio"
	"fmt"
	"io"
)

type StringMap struct {
	Value map[string]string
	Sep   byte
	Comma byte
	Delim byte
}

func (step *StringMap) Write(writer io.Writer) error {
	for key, value := range step.Value {
		_, err := fmt.Fprintf(writer, "%s%s%s%s", key, string(step.Sep), value, string(step.Comma))
		if err != nil {
			return err
		}
	}
	_, err := writer.Write([]byte{step.Delim})
	return err
}

func (step *StringMap) Read(reader *bufio.Reader) error {
	for {
		peek, err := reader.Peek(1)
		if err == nil && peek[0] == step.Delim {
			reader.Discard(1)
			break
		}

		key, err := reader.ReadString(step.Sep)
		if err != nil {
			return err
		}

		value, err := reader.ReadString(step.Comma)
		if err != nil {
			return err
		}

		step.Value[key[:len(key)-1]] = value[:len(value)-1]
	}

	return nil
}
