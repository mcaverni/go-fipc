package base

import "fmt"

type Value interface {
	FromBytes([]byte) (Value, error)
	ToBytes() ([]byte, error)
}

type ValueInt int32
type ValueFloat float32 // TODO: implement type functions
type ValueString string // TODO: implement type functions
type ValueBool bool     // TODO: implement type functions
type ValueChar rune     // TODO: implement type functions

func (v ValueInt) FromBytes(b []byte) (Value, error) {
	// TODO: implement
	return ValueInt(0), fmt.Errorf("not implemented")
}

func (v ValueInt) ToBytes() (b []byte, e error) {
	// TODO: implement
	return make([]byte, 10), fmt.Errorf("not implemented")
}

func BytesToValue(b []byte) (Value, error) {
	// TODO: implement
	return Value(nil), fmt.Errorf("not implemented")
}
