package base

type Value interface {
	FromBytes([]byte) Value
	ToBytes() []byte
}

type ValueInt int32

func (v ValueInt) FromBytes(b []byte) Value {
	return ValueInt(0)
}

func (v ValueInt) ToBytes() []byte {
	return make([]byte, 10)
}
