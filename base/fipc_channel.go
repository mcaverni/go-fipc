package base

type FIPC_Channel struct {
	Id       FIPC_Identifier
	Name     string
	Value    Value
	Max      Value
	Min      Value
	Default  Value
	Unit     Unit
	Visible  Condition
	Disabled Condition
	Change   Condition
}

func (c FIPC_Channel) toBytes() []byte {
	// TODO: implement
	return make([]byte, 10)
}

func (c FIPC_Channel) fromBytes(b []byte) FIPC_Channel {
	// TODO: implement
	return FIPC_Channel{}
}

func (c FIPC_Channel) toInt() int {
	// TODO: implement
	return 0
}

func (c FIPC_Channel) toFloat() float32 {
	// TODO: implement
	return 0.0
}

func (c FIPC_Channel) toBool() bool {
	// TODO: implement
	return false
}

func (c FIPC_Channel) toString() string {
	// TODO: implement
	return ""
}
