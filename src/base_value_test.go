package gofipc

import (
	"testing"
)

func TestInteger(t *testing.T) {
	v := ValueInt(0x12345678)
	buf, err := v.MarshalBinary()
	if err != nil {
		t.Error(err)
	}
	if len(buf) != 4 {
		t.Errorf("Expected 4 bytes, got %d", len(buf))
	}

	var v2 ValueInt
	err = v2.UnmarshalBinary(buf)
	if err != nil {
		t.Error(err)
	}
	if v != v2 {
		t.Errorf("Expected %d, got %d", v, v2)
	}
}

func TestFloat(t *testing.T) {
	v := ValueFloat(3.14159)
	buf, err := v.MarshalBinary()
	if err != nil {
		t.Error(err)
	}
	if len(buf) != 4 {
		t.Errorf("Expected 4 bytes, got %d", len(buf))
	}

	var v2 ValueFloat
	err = v2.UnmarshalBinary(buf)
	if err != nil {
		t.Error(err)
	}
	if v != v2 {
		t.Errorf("Expected %f, got %f", v, v2)
	}
}

func TestString(t *testing.T) {
	v := ValueString("Hello, World!")
	buf, err := v.MarshalBinary()
	if err != nil {
		t.Error(err)
	}
	if len(buf) != len(v) {
		t.Errorf("Expected %d bytes, got %d", len(v), len(buf))
	}
	var v2 ValueString
	err = v2.UnmarshalBinary(buf)
	if err != nil {
		t.Error(err)
	}
	if v != v2 {
		t.Errorf("Expected %s, got %s", v, v2)
	}
}

func TestBool(t *testing.T) {
	v := ValueBool(true)
	buf, err := v.MarshalBinary()
	if err != nil {
		t.Error(err)
	}
	if len(buf) != 1 {
		t.Errorf("Expected 1 byte, got %d", len(buf))
	}

	var v2 ValueBool
	err = v2.UnmarshalBinary(buf)
	if err != nil {
		t.Error(err)
	}
	if v != v2 {
		t.Errorf("Expected %t, got %t", v, v2)
	}
}

func TestChar(t *testing.T) {
	v := ValueChar('A')
	buf, err := v.MarshalBinary()
	if err != nil {
		t.Error(err)
	}
	if len(buf) != 1 {
		t.Errorf("Expected 4 bytes, got %d", len(buf))
	}

	var v2 ValueChar
	err = v2.UnmarshalBinary(buf)
	if err != nil {
		t.Error(err)
	}
	if v != v2 {
		t.Errorf("Expected %c, got %c", v, v2)
	}
}
