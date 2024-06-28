package base

import "fmt"

type Channel struct {
	Id       ChannelId
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

func (c Channel) ToBytes() ([]byte, error) {
	// TODO: implement
	return make([]byte, 10), fmt.Errorf("not implemented")
}

func (c Channel) ToString() (string, error) {
	// TODO: implement
	return "", fmt.Errorf("not implemented")
}

func BytesToChannel(b []byte) (Channel, error) {
	// TODO: implement
	return Channel{}, fmt.Errorf("not implemented")
}
