package gomethius

import (
	"bytes"
	"encoding/binary"
)

type Value interface {
	Encode() []byte
	Data() interface{}
}

type Values struct {
	values []Value
}

func NewValues(values ...Value) *Values {
	out := new(Values)
	out.values = values
	return out
}

func (v *Values) Get(idx int) Value {
	return v.values[idx]
}


type StringValue struct {
	data string
}

func NewStringValue(d string) *StringValue {
	out := new(StringValue)
	out.data = d
	return out;
}

func (v *StringValue) Encode() []byte {
	return []byte(v.data)
}

func (v *StringValue) Data() interface{} {
	return v.data
}

type Float32Value struct {
	data float32
}

func NewFloat32Value(d float32) *Float32Value {
	out := new(Float32Value)
	out.data = d
	return out;
}

func (v *Float32Value) Encode() []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, v.data)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

func (v *Float32Value) Data() interface{} {
	return v.data
}

type Float64Value struct {
	data float64
}

func NewFloat64Value(d float64) *Float64Value {
	out := new(Float64Value)
	out.data = d
	return out;
}

func (v *Float64Value) Encode() []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, v.data)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

func (v *Float64Value) Data() interface{} {
	return v.data
}

