package base

type Value interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary([]byte) error
}

type ValueInt int32
type ValueFloat float32
type ValueString string
type ValueBool bool
type ValueChar rune

func (v ValueInt) MarshalBinary() ([]byte, error) {
	return []byte{byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)}, nil
}

func (v *ValueInt) UnmarshalBinary(b []byte) error {
	*v = ValueInt(int32(b[0])<<24 | int32(b[1])<<16 | int32(b[2])<<8 | int32(b[3]))
	return nil
}
