package base

import "fmt"

type ChannelType uint8

const CHANNEL_ANALOG_INPUT = ChannelType(0)
const CHANNEL_ANALOG_OUTPUT = ChannelType(1)
const CHANNEL_COMMAND = ChannelType(2)
const CHANNEL_EVENT = ChannelType(3)
const CHANNEL_SETUP = ChannelType(4)
const CHANNEL_DIGITAL_INPUT = ChannelType(5)
const CHANNEL_DIGITAL_OUTPUT = ChannelType(6)
const CHANNEL_STATUS = ChannelType(7)
const CHANNEL_INTERNAL = ChannelType(50)
const CHANNEL_NONE = ChannelType(255)

func TypeToString(t ChannelType) string {
	switch t {
	case CHANNEL_ANALOG_INPUT:
		return "anin"
	case CHANNEL_ANALOG_OUTPUT:
		return "anout"
	case CHANNEL_COMMAND:
		return "command"
	case CHANNEL_EVENT:
		return "alarm"
	case CHANNEL_SETUP:
		return "setup"
	case CHANNEL_DIGITAL_INPUT:
		return "digin"
	case CHANNEL_DIGITAL_OUTPUT:
		return "digout"
	case CHANNEL_STATUS:
		return "status"
	case CHANNEL_INTERNAL:
		return "internal"
	case CHANNEL_NONE:
		return "none"
	}
	return ""
}

func StringToType(s string) ChannelType {
	switch s {
	case "anin":
		return CHANNEL_ANALOG_INPUT
	case "anout":
		return CHANNEL_ANALOG_OUTPUT
	case "command":
		return CHANNEL_COMMAND
	case "alarm":
		return CHANNEL_EVENT
	case "SETUP":
		return CHANNEL_SETUP
	case "digin":
		return CHANNEL_DIGITAL_INPUT
	case "digout":
		return CHANNEL_DIGITAL_OUTPUT
	case "status":
		return CHANNEL_STATUS
	case "internal":
		return CHANNEL_INTERNAL
	case "none":
		return CHANNEL_NONE
	}
	return CHANNEL_NONE
}

type FIPC_Identifier struct {
	Type    ChannelType
	Address int32 // slave address
	Channel int32 // channel ID
}

func (i FIPC_Identifier) ToBytes() []byte {
	// TODO: implement
	return make([]byte, 10)
}

func (i FIPC_Identifier) FromBytes(b []byte) FIPC_Identifier {
	// TODO: implement
	return FIPC_Identifier{}
}

func (i FIPC_Identifier) ToString() string {
	return fmt.Sprintf("%d:%s:%d", i.Address, TypeToString(i.Type), i.Channel)
}

func (i FIPC_Identifier) FromString(s string) {
	var tmp string
	fmt.Sscanf(s, "%d:%s:%d", &i.Address, tmp, &i.Channel)
	i.Type = StringToType(tmp)
}
