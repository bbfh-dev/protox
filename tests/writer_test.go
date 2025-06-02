package protox_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/bbfh-dev/protox"
	"gotest.tools/assert"
)

func TestWriter(test *testing.T) {
	testCases := []struct {
		Call      func(writer io.Writer) error
		ExpectErr error
		Expect    string
	}{
		{
			Call: func(writer io.Writer) error {
				return protox.NewWriter(writer).
					Strings("example string").
					Bytes('|').
					Ints(42069).
					Bytes('|').
					Bools(true, false).
					Bytes('|').
					Floats(123.456).
					Bytes('|').
					Array([]string{"a", "b", "c"}, []byte(", "), func(i int, value string) []byte {
						return []byte(value)
					}).
					Complete()
			},
			ExpectErr: nil,
			Expect:    "example string|42069|truefalse|123.456|a, b, c",
		},
	}

	for _, testCase := range testCases {
		test.Run("", func(test *testing.T) {
			var buffer bytes.Buffer
			err := testCase.Call(&buffer)
			if testCase.ExpectErr == nil {
				assert.NilError(test, err)
			} else {
				assert.ErrorContains(test, err, testCase.ExpectErr.Error())
			}
			assert.DeepEqual(test, buffer.String(), testCase.Expect)
		})
	}
}
