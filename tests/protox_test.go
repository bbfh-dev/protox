package protox_test

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/bbfh-dev/protox"
	"gotest.tools/assert"
)

type Example struct {
	Byte        byte
	BytesDelim  []byte
	String      string
	StringDelim string
	Int         int64
	IntDelim    int64
	FloatDelim  float64
	StringMap   map[string]string
	StringArr   []string
}

func (example *Example) Protox() *protox.Processor {
	return protox.New().
		ThenByte(&example.Byte).
		ThenBytesDelim(&example.BytesDelim, '\x00').
		ThenString(&example.String, 3).
		ThenStringDelim(&example.StringDelim, '\x00').
		ThenInt(&example.Int).
		ThenIntDelim(&example.IntDelim, '\x00').
		ThenFloatDelim(&example.FloatDelim, '\x00').
		ThenStringMap(example.StringMap, '=', '\x00', '\x00').
		ThenStringArray(&example.StringArr, ',', ';').
		Build()
}

var example = &Example{
	Byte:        'g',
	BytesDelim:  []byte{'a', 'b', 'c'},
	String:      "hi!",
	StringDelim: "Hello World!",
	Int:         1337,
	IntDelim:    42069,
	FloatDelim:  420.69,
	StringMap:   map[string]string{},
	StringArr:   []string{"abc", "xyz", "123"},
}

var exampleString = "g" + "abc\x00" + "hi!" + "Hello World!\x00" + "9\x05\x00\x00\x00\x00\x00\x00" + "42069\x00" + "420.69\x00\x00abc,xyz,123,;"

func TestWrite(test *testing.T) {
	var buffer bytes.Buffer
	assert.NilError(test, example.Protox().Write(&buffer))
	test.Logf("Result: %q\n", buffer.String())
	assert.DeepEqual(test, buffer.String(), exampleString)
}

func TestRead(test *testing.T) {
	var buffer bytes.Buffer
	assert.NilError(test, example.Protox().Write(&buffer))

	instance := &Example{
		StringMap: map[string]string{},
		StringArr: []string{},
	}
	test.Logf("Before: %#v\n", instance)
	assert.NilError(test, instance.Protox().Read(bufio.NewReader(&buffer)))
	test.Logf("After: %#v\n", instance)
	assert.DeepEqual(test, instance.Byte, example.Byte)
	assert.DeepEqual(test, instance.BytesDelim, example.BytesDelim)
	assert.DeepEqual(test, instance.String, example.String)
	assert.DeepEqual(test, instance.StringDelim, example.StringDelim)
	assert.DeepEqual(test, instance.Int, example.Int)
	assert.DeepEqual(test, instance.IntDelim, example.IntDelim)
	assert.DeepEqual(test, instance.FloatDelim, example.FloatDelim)
	assert.DeepEqual(test, instance.StringMap, example.StringMap)
	assert.DeepEqual(test, instance.StringArr, example.StringArr)
}
