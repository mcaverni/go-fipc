package base

import (
	"testing"
)

func TestStruct(t *testing.T) {
	// test equivalcence between these two
	id_value := ChannelId{CHANNEL_ANALOG_INPUT, 1, 15}

	if id_value.Type != CHANNEL_ANALOG_INPUT || id_value.Address != 1 || id_value.Channel != 15 {
		t.Error("ChannelId struct failed")
	}
}

func TestStringsAnin(t *testing.T) {
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

func TestStringsAnout(t *testing.T) {
	// test equivalcence between these two
	const id_string string = "1 anout 15"
	id_value := ChannelId{CHANNEL_ANALOG_OUTPUT, 1, 15}

	s, e := id_value.ToString()
	if e != nil || s != id_string {
		t.Error("ChannelId.ToString() failed", e)
	}

	c, e := StringToChannelId(id_string)
	if e != nil || c != id_value {
		t.Error("StringToChannelId() failed:", e)
	}
}

func TestStringsCommand(t *testing.T) {
	// test equivalcence between these two
	const id_string string = "1 command 15"
	id_value := ChannelId{CHANNEL_COMMAND, 1, 15}

	s, e := id_value.ToString()
	if e != nil || s != id_string {
		t.Error("ChannelId.ToString() failed", e)
	}

	c, e := StringToChannelId(id_string)
	if e != nil || c != id_value {
		t.Error("StringToChannelId() failed:", e)
	}
}

func TestStringsAlarm(t *testing.T) {
	// test equivalcence between these two
	const id_string string = "1 alarm 15"
	id_value := ChannelId{CHANNEL_EVENT, 1, 15}

	s, e := id_value.ToString()
	if e != nil || s != id_string {
		t.Error("ChannelId.ToString() failed", e)
	}

	c, e := StringToChannelId(id_string)
	if e != nil || c != id_value {
		t.Error("StringToChannelId() failed:", e)
	}
}

func TestStringsSetup(t *testing.T) {
	// test equivalcence between these two
	const id_string string = "1 setup 15"
	id_value := ChannelId{CHANNEL_SETUP, 1, 15}

	s, e := id_value.ToString()
	if e != nil || s != id_string {
		t.Error("ChannelId.ToString() failed", e)
	}

	c, e := StringToChannelId(id_string)
	if e != nil || c != id_value {
		t.Error("StringToChannelId() failed:", e)
	}
}

func TestStringsDigin(t *testing.T) {
	// test equivalcence between these two
	const id_string string = "1 digin 15"
	id_value := ChannelId{CHANNEL_DIGITAL_INPUT, 1, 15}

	s, e := id_value.ToString()
	if e != nil || s != id_string {
		t.Error("ChannelId.ToString() failed", e)
	}

	c, e := StringToChannelId(id_string)
	if e != nil || c != id_value {
		t.Error("StringToChannelId() failed:", e)
	}
}

func TestStringsDigout(t *testing.T) {
	// test equivalcence between these two
	const id_string string = "1 digout 15"
	id_value := ChannelId{CHANNEL_DIGITAL_OUTPUT, 1, 15}

	s, e := id_value.ToString()
	if e != nil || s != id_string {
		t.Error("ChannelId.ToString() failed", e)
	}

	c, e := StringToChannelId(id_string)
	if e != nil || c != id_value {
		t.Error("StringToChannelId() failed:", e)
	}
}

func TestStringsStatus(t *testing.T) {
	// test equivalcence between these two
	const id_string string = "1 status 15"
	id_value := ChannelId{CHANNEL_STATUS, 1, 15}

	s, e := id_value.ToString()
	if e != nil || s != id_string {
		t.Error("ChannelId.ToString() failed", e)
	}

	c, e := StringToChannelId(id_string)
	if e != nil || c != id_value {
		t.Error("StringToChannelId() failed:", e)
	}
}

func TestStringsInternal(t *testing.T) {
	// test equivalcence between these two
	const id_string string = "1 internal 15"
	id_value := ChannelId{CHANNEL_INTERNAL, 1, 15}

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

	e := c.UnmarshalBinary([]byte{0, 0, 0, 0})
	if e == nil {
		t.Error("BytesToChannelId() failed to detect invalid data")
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

func TestStringErrors(t *testing.T) {
	// test invalid channel type
	_, e := StringToChannelId("1 invalid 15")
	if e == nil {
		t.Error("StringToChannelId() failed to detect invalid channel type")
	}

	_, e = StringToChannelId("1 anin")
	if e == nil {
		t.Error("StringToChannelId() failed to detect missing channel ID")
	}

	_, e = StringToChannelId("1 anin fifteen")
	if e == nil {
		t.Error("StringToChannelId() failed to detect invalid channel ID")
	}

	var c ChannelId
	// c, e = StringToChannelId("1 anin 15 15")
	// if e == nil {
	// 	cs, _ := c.ToString()
	// 	t.Error("StringToChannelId() failed to detect extra data", cs)
	// }

	c = ChannelId{ChannelType(85), 1, 15}
	_, e = c.ToString()
	if e == nil {
		t.Error("ChannelId.ToString() failed to detect invalid channel type")
	}

	_, e = ChannelTypeToString(ChannelType(85))
	if e == nil {
		t.Error("ChannelTypeToString() failed to detect invalid channel type")
	}
}
