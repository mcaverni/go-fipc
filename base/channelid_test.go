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
	// const id_bytes [6]byte = [6]byte{0, 1, 15, 0, 0, 0}
	id_value := ChannelId{CHANNEL_ANALOG_INPUT, 1, 15}

	b, e := id_value.ToBytes()
	if e != nil {
		t.Error("ChannelId.ToBytes() failed", e)
	}

	c, e := BytesToChannelId(b)
	if e != nil || c != id_value {
		t.Error("BytesToChannelId() failed:", e)
	}
}
