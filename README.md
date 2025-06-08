# Protox

Go library for (de-)serializing custom protocols

# Usage

Use `protox.New().<...>.Build()` to define a processor that uses a specific format for writing/reading. You can call `Read` or `Write` methods on the resulting `*protox.Processor` to (de-)serialize data.

Here's an example using a struct, it uses all allowed variable formats:

```go
import "github.com/bbfh-dev/protox"

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

func write(buffer *bytes.Buffer) {
    err := example.Protox().Write(buffer)
    fmt.Println(buffer.String()) // "gabc\x00hi!Hello World!\x009\x05\x00\x00\x00\x00\x00\x0042069\x00420.69\x00a=(1)\x00b=(2)\x00c=(3)\x00\x00"
}

func read(buffer *bytes.Buffer) {
    instance := &Example{}
    err := instance.Protox().Read(bufio.NewReader(buffer))
    fmt.Prinln(instance) // &protox_test.Example{Byte:0x67, BytesDelim:[]uint8{0x61, 0x62, 0x63}, String:"hi!", StringDelim:"Hello World!", Int:1337, IntDelim:42069, FloatDelim:420.69, StringMap:map[string]string{"a":"(1)", "b":"(2)", "c":"(3)"}}
}
```
