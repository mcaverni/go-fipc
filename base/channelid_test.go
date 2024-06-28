package base

import (
	"testing"
)

func TestStrings(t *testing.T) {
	// test equivalcence between these two
	const id_string string = "1 anin 15"
	id_value := ChannelId{CHANNEL_ANALOG_INPUT, 1, 15}

	s, e := id_value.ToString()
	if e != nil || s != id_string {
		t.Error("ChannelId.ToString() failed", e)
	}

	c, e := StringToChannelId(id_string)
	if e != nil || c != id_value {
		t.Error("StringToChannelId() failed:", e)
	}
}

func TestBytes(t *testing.T) {
	// test equivalcence between these two
	chan_id := ChannelId{CHANNEL_ANALOG_INPUT, 1, 15}

	byte_array, err := chan_id.MarshalBinary()
	if err != nil {
		t.Error("ChannelId.ToBytes() failed", err)
	}

	c := ChannelId{}
	err = c.UnmarshalBinary(byte_array)
	if err != nil || c != chan_id {
		t.Error("BytesToChannelId() failed:", err)
	}
}

func TestInvalid(t *testing.T) {
	// test invalid channel type
	nochan := NoChannelId()
	if nochan.IsValid() {
		t.Error("NoChannelId.IsValid() failed")
	}

	yeschan := ChannelId{CHANNEL_ANALOG_INPUT, 1, 15}
	if !yeschan.IsValid() {
		t.Error("ChannelId.IsValid() failed")
	}
}
