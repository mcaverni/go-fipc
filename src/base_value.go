package gofipc

import (
	"fmt"
	"math"
)

type Value interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary([]byte) error
	// TODO: implement those
	ToInt() (ValueInt, error)
	ToFloat() (ValueFloat, error)
	ToString() (ValueString, error)
	ToBool() (ValueBool, error)
	ToChar() (ValueChar, error)
}

type ValueInt int32
type ValueFloat float32
type ValueString string
type ValueBool bool
type ValueChar byte

func (v ValueInt) MarshalBinary() ([]byte, error) {
	return []byte{byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)}, nil
}

func (v *ValueInt) UnmarshalBinary(b []byte) error {
	if len(b) != 4 {
		return fmt.Errorf("invalid buffer length")
	}
	*v = ValueInt(int32(b[0])<<24 | int32(b[1])<<16 | int32(b[2])<<8 | int32(b[3]))
	return nil
}

func (v ValueFloat) MarshalBinary() ([]byte, error) {
	i := math.Float32bits(float32(v))
	return []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}, nil
}

func (v *ValueFloat) UnmarshalBinary(b []byte) error {
	if len(b) != 4 {
		return fmt.Errorf("invalid buffer length")
	}
	i := uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
	*v = ValueFloat(math.Float32frombits(i))
	return nil
}

func (v ValueString) MarshalBinary() ([]byte, error) {
	return []byte(v), nil
}

func (v *ValueString) UnmarshalBinary(b []byte) error {
	*v = ValueString(b)
	return nil
}

func (v ValueBool) MarshalBinary() ([]byte, error) {
	if v {
		return []byte{1}, nil
	}
	return []byte{0}, nil
}

func (v *ValueBool) UnmarshalBinary(b []byte) error {
	if len(b) != 1 {
		return fmt.Errorf("invalid buffer length")
	}
	*v = b[0] != 0
	return nil
}

func (v ValueChar) MarshalBinary() ([]byte, error) {
	return []byte{byte(v)}, nil
}

func (v *ValueChar) UnmarshalBinary(b []byte) error {
	if len(b) != 1 {
		return fmt.Errorf("invalid buffer length")
	}
	*v = ValueChar(byte(b[0]))
	return nil
}
