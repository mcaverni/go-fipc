package base

import "fmt"

type ChannelType uint8

const CHANNEL_ANALOG_INPUT ChannelType = ChannelType(0)
const CHANNEL_ANALOG_OUTPUT ChannelType = ChannelType(1)
const CHANNEL_COMMAND ChannelType = ChannelType(2)
const CHANNEL_EVENT ChannelType = ChannelType(3)
const CHANNEL_SETUP ChannelType = ChannelType(4)
const CHANNEL_DIGITAL_INPUT ChannelType = ChannelType(5)
const CHANNEL_DIGITAL_OUTPUT ChannelType = ChannelType(6)
const CHANNEL_STATUS ChannelType = ChannelType(7)
const CHANNEL_INTERNAL ChannelType = ChannelType(50)
const CHANNEL_NONE ChannelType = ChannelType(255)

func ChannelTypeToString(t ChannelType) (string, error) {
	switch t {
	case CHANNEL_ANALOG_INPUT:
		return "anin", nil
	case CHANNEL_ANALOG_OUTPUT:
		return "anout", nil
	case CHANNEL_COMMAND:
		return "command", nil
	case CHANNEL_EVENT:
		return "alarm", nil
	case CHANNEL_SETUP:
		return "setup", nil
	case CHANNEL_DIGITAL_INPUT:
		return "digin", nil
	case CHANNEL_DIGITAL_OUTPUT:
		return "digout", nil
	case CHANNEL_STATUS:
		return "status", nil
	case CHANNEL_INTERNAL:
		return "internal", nil
	}
	return "none", fmt.Errorf("invalid channel type: %d", t)
}

func StringToChannelType(s string) (ChannelType, error) {
	switch s {
	case "anin":
		return CHANNEL_ANALOG_INPUT, nil
	case "anout":
		return CHANNEL_ANALOG_OUTPUT, nil
	case "command":
		return CHANNEL_COMMAND, nil
	case "alarm":
		return CHANNEL_EVENT, nil
	case "SETUP":
		return CHANNEL_SETUP, nil
	case "digin":
		return CHANNEL_DIGITAL_INPUT, nil
	case "digout":
		return CHANNEL_DIGITAL_OUTPUT, nil
	case "status":
		return CHANNEL_STATUS, nil
	case "internal":
		return CHANNEL_INTERNAL, nil
	}
	return CHANNEL_NONE, fmt.Errorf("invalid channel type: %s", s)
}

type ChannelId struct {
	Type    ChannelType
	Address int32 // slave address
	Channel int32 // channel ID
}

func (i ChannelId) ToBytes() ([]byte, error) {
	// TODO: implement
	return make([]byte, 10), fmt.Errorf("not implemented")
}

func (i ChannelId) ToString() (string, error) {
	s, e := ChannelTypeToString(i.Type)
	if e != nil {
		return "", e
	}
	return fmt.Sprintf("%d %s %d", i.Address, s, i.Channel), nil
}

func BytesToChannelId(b []byte) (ChannelId, error) {
	// TODO: implement
	return ChannelId{}, fmt.Errorf("not implemented")
}

func StringToChannelId(s string) (ChannelId, error) {
	var tmp string
	var ad, ch int32
	n, err := fmt.Sscanf(s, "%d %s %d", &ad, &tmp, &ch)
	if n != 3 || err != nil {
		return ChannelId{}, err
	}
	ty, err := StringToChannelType(tmp)
	if err != nil {
		return ChannelId{}, err
	}
	return ChannelId{ty, ad, ch}, nil
}
