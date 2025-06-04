package protox

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"github.com/bbfh-dev/protox/internal"
)

type Processor struct {
	steps []internal.Step
}

// Returns on the first error encountered
func (processor *Processor) Read(reader *bufio.Reader) error {
	for i, step := range processor.steps {
		err := step.Read(reader)
		if err != nil {
			return fmt.Errorf("protox.Processor.step(%d): %w", i, err)
		}
	}

	return nil
}

// Continues to read even if an error is encountered, returns all errors combined
func (processor *Processor) ForceRead(reader *bufio.Reader) error {
	errs := []error{}

	for i, step := range processor.steps {
		err := step.Read(reader)
		if err != nil {
			errs = append(errs, fmt.Errorf("protox.Processor(reading).step(%d): %w", i, err))
		}
	}

	return errors.Join(errs...)
}

// Returns on the first error encountered
func (processor *Processor) Write(writer io.Writer) error {
	for i, step := range processor.steps {
		err := step.Write(writer)
		if err != nil {
			return fmt.Errorf("protox.Processor(writing).step(%d): %w", i, err)
		}
	}

	return nil
}

// Continues to write even if an error is encountered, returns all errors combined
func (processor *Processor) ForceWrite(writer io.Writer) error {
	errs := []error{}

	for i, step := range processor.steps {
		err := step.Write(writer)
		if err != nil {
			errs = append(errs, fmt.Errorf("protox.Processor(writing).step(%d): %w", i, err))
		}
	}

	return errors.Join(errs...)
}
