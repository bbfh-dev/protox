package protox

import "github.com/bbfh-dev/protox/internal"

type Builder struct {
	steps []internal.Step
}

func New() *Builder {
	return &Builder{
		steps: []internal.Step{},
	}
}

// Build the Protox processor
func (build *Builder) Build() *Processor {
	return &Processor{
		steps: build.steps,
	}
}

func (builder *Builder) ThenByte(ref *byte) *Builder {
	builder.steps = append(builder.steps, &internal.Byte{
		Ref: ref,
	})
	return builder
}

func (builder *Builder) ThenBytesDelim(value *[]byte, delim byte) *Builder {
	builder.steps = append(builder.steps, &internal.ByteDelim{
		Value: value,
		Delim: delim,
	})
	return builder
}

func (builder *Builder) ThenString(ref *string, size int) *Builder {
	builder.steps = append(builder.steps, &internal.String{
		Ref:  ref,
		Size: size,
	})
	return builder
}

func (builder *Builder) ThenStringDelim(ref *string, delim byte) *Builder {
	builder.steps = append(builder.steps, &internal.StringDelim{
		Ref:   ref,
		Delim: delim,
	})
	return builder
}

func (builder *Builder) ThenInt(ref *int64) *Builder {
	builder.steps = append(builder.steps, &internal.Int{
		Ref: ref,
	})
	return builder
}

func (builder *Builder) ThenIntDelim(ref *int64, delim byte) *Builder {
	builder.steps = append(builder.steps, &internal.IntDelim{
		Ref:   ref,
		Delim: delim,
	})
	return builder
}

func (builder *Builder) ThenFloatDelim(ref *float64, delim byte) *Builder {
	builder.steps = append(builder.steps, &internal.FloatDelim{
		Ref:   ref,
		Delim: delim,
	})
	return builder
}

func (builder *Builder) ThenStringMap(value map[string]string, sep, comma, delim byte) *Builder {
	builder.steps = append(builder.steps, &internal.StringMap{
		Value: value,
		Sep:   sep,
		Comma: comma,
		Delim: delim,
	})
	return builder
}
