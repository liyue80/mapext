package mapext

import (
	"encoding/json"
	"strings"
	"testing"
)

type Struct struct {
	BoolValue   bool
	IntValue    int
	Int8Value   int8
	Int16Value  int16
	Int32Value  int32
	Int64Value  int64
	UintValue   uint
	Uint8Value  uint8
	Uint16Value uint16
	Uint32Value uint32
	Uint64Value uint64
	StringValue string
}

func Test_GetStringValue_1(t *testing.T) {
	b := []byte(`{"Name":"Wednesday"}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	if name, err := GetStringValue(m, "Name"); err == nil {
		if strings.Compare(name, "Wednesday") == 0 {
			t.Log("Pass")
		} else {
			t.Error("Error")
		}
	} else {
		t.Error("Error")
	}
}
