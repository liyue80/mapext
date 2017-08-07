package mapext

import (
	"encoding/json"
	"strings"
	"testing"
)

//type Struct struct {
//	BoolValue   bool
//	IntValue    int
//	Int8Value   int8
//	Int16Value  int16
//	Int32Value  int32
//	Int64Value  int64
//	UintValue   uint
//	Uint8Value  uint8
//	Uint16Value uint16
//	Uint32Value uint32
//	Uint64Value uint64
//	StringValue string
//}

func TestGetStringValue_1(t *testing.T) {
	b := []byte(`{"Name":"Wednesday"}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	if name, err := GetStringValue(m, "Name"); err == nil {
		if strings.Compare(name, "Wednesday") != 0 {
			t.Error("func GetStringValue return a wrong string value")
		}
	} else {
		t.Error("func GetStringValue returns with error")
	}
}

func TestGetStringValue_NoKey(t *testing.T) {
	b := []byte(`{"Name":"Wednesday"}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	if _, err := GetStringValue(m, "Key-not-exists"); err == nil {
		t.Error("func GetStringValue should return with error")
	}
}

func TestGetBoolValue(t *testing.T) {
	b := []byte(`{"t1":true,"t2":"True","t3":"TRUE","t4":"TrUe","t5":1,"t6":"1","f1":false,"int0":0}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	var value bool
	var err error

	if value, err = GetBoolValue(m, "t1"); err == nil {
		if value {
			t.Log("Get value true")
		} else {
			t.Error("GetBoolValue returns a wrong value (except for true")
		}
	} else {
		t.Error("func GetBoolValue should return no error")
	}

	if value, err = GetBoolValue(m, "t2"); err == nil {
		if value {
			t.Log("Get value true")
		} else {
			t.Error("GetBoolValue returns a wrong value (except for true")
		}
	} else {
		t.Error("func GetBoolValue should return no error")
	}

	if value, err = GetBoolValue(m, "t3"); err == nil {
		if value {
			t.Log("Get value true")
		} else {
			t.Error("GetBoolValue returns a wrong value (except for true")
		}
	} else {
		t.Error("func GetBoolValue should return no error")
	}

	if value, err = GetBoolValue(m, "t4"); err != nil {
		if strings.Index(err.Error(), "invalid syntax") != -1 {
			t.Log("Get value true")
		} else {
			t.Errorf("func GetBoolValue return incorrect error type")
		}
	} else {
		t.Error("func GetBoolValue should return with error")
	}

	if value, err = GetBoolValue(m, "t5"); err == nil {
		if value {
			t.Log("Get value true")
		} else {
			t.Error("GetBoolValue returns a wrong value (except for true")
		}
	} else {
		t.Error("func GetBoolValue should return no error")
	}

	if value, err = GetBoolValue(m, "t6"); err == nil {
		if value {
			t.Log("Get value true")
		} else {
			t.Error("GetBoolValue returns a wrong value (except for true")
		}
	} else {
		t.Error("func GetBoolValue should return no error")
	}

	if value, err = GetBoolValue(m, "f1"); err == nil {
		if value {
			t.Error("GetBoolValue returns a wrong value (except for false")
		} else {
			t.Log("Get value false")
		}
	} else {
		t.Error("func GetBoolValue should return no error")
	}

	if value, err = GetBoolValue(m, "int0"); err == nil {
		if value {
			t.Error("GetBoolValue returns a wrong value (except for false")
		} else {
			t.Log("Get value false")
		}
	} else {
		t.Error("func GetBoolValue should return no error")
	}
}

func TestGetInt8(t *testing.T) {
	b := []byte(`{"zero":0,"zero-s":"0","one":1,"one-s":"1","v127":127,"v127-s":"127","v128":128,"v128-s":"128","n128":-128,"n128-s":"-128","n129":-129,"n129-s":"-129"}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	var value int8
	var err error
	if value, err = GetInt8Value(m, "zero"); err == nil {
		if value != 0 {
			t.Error("func GetInt8Value return a wrong value (expected for 0)")
		} else {
			t.Log("Get value 0")
		}
	} else {
		t.Error("func GetInt8Value should return no error")
	}

	if value, err = GetInt8Value(m, "zero-s"); err == nil {
		if value != 0 {
			t.Error("func GetInt8Value return a wrong value (expected for 0)")
		} else {
			t.Log("Get value 0")
		}
	} else {
		t.Error("func GetInt8Value should return no error")
	}

	if value, err = GetInt8Value(m, "one"); err == nil {
		if value != 1 {
			t.Error("func GetInt8Value return a wrong value (expected for 1)")
		} else {
			t.Log("Get value 1")
		}
	} else {
		t.Error("func GetInt8Value should return no error")
	}

	if value, err = GetInt8Value(m, "one-s"); err == nil {
		if value != 1 {
			t.Error("func GetInt8Value return a wrong value (expected for 1)")
		} else {
			t.Log("Get value 1")
		}
	} else {
		t.Error("func GetInt8Value should return no error")
	}

	if value, err = GetInt8Value(m, "v127"); err == nil {
		if value != 127 {
			t.Error("func GetInt8Value return a wrong value (expected for 127)")
		} else {
			t.Log("Get value 127")
		}
	} else {
		t.Error("func GetInt8Value should return no error")
	}

	if value, err = GetInt8Value(m, "v127-s"); err == nil {
		if value != 127 {
			t.Error("func GetInt8Value return a wrong value (expected for 127)")
		} else {
			t.Log("Get value 127")
		}
	} else {
		t.Error("func GetInt8Value should return no error")
	}

	if _, err = GetInt8Value(m, "v128"); err != nil {
		if strings.Index(err.Error(), "value out of range") != -1 {
			t.Log("Get value 128")
		} else {
			t.Errorf("func GetInt8Value return incorrect error type")
		}
	} else {
		t.Error("func GetInt8Value should return with error")
	}

	if _, err = GetInt8Value(m, "v128-s"); err != nil {
		if strings.Index(err.Error(), "value out of range") != -1 {
			t.Log("Get value 128")
		} else {
			t.Error("func GetInt8Value return incorrect error type")
		}
	} else {
		t.Error("func GetInt8Value should return with error")
	}

	if value, err = GetInt8Value(m, "n128"); err == nil {
		if value != -128 {
			t.Error("func GetInt8Value return a wrong value (expected for -128)")
		} else {
			t.Log("Get value -128")
		}
	} else {
		t.Error("func GetInt8Value should return no error")
	}

	if value, err = GetInt8Value(m, "n128-s"); err == nil {
		if value != -128 {
			t.Error("func GetInt8Value return a wrong value (expected for -128)")
		} else {
			t.Log("Get value -128")
		}
	} else {
		t.Error("func GetInt8Value should return no error")
	}

	if _, err = GetInt8Value(m, "n129"); err != nil {
		if strings.Index(err.Error(), "value out of range") != -1 {
			t.Log("Get value -129")
		} else {
			t.Errorf("func GetInt8Value return incorrect error type")
		}
	} else {
		t.Error("func GetInt8Value should return with error")
	}

	if _, err = GetInt8Value(m, "n129-s"); err != nil {
		if strings.Index(err.Error(), "value out of range") != -1 {
			t.Log("Get value -129")
		} else {
			t.Error("func GetInt8Value return incorrect error type")
		}
	} else {
		t.Error("func GetInt8Value should return with error")
	}
}

func TestGetInt16Value(t *testing.T) {
	b := []byte(`{"one":1}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	var value int16
	var err error
	if value, err = GetInt16Value(m, "one"); err == nil {
		if value != 1 {
			t.Error("func GetInt16Value return a wrong value (expected for 1)")
		} else {
			t.Log("Get value 1")
		}
	} else {
		t.Error("func GetInt16Value should return no error")
	}
}

func TestGetInt32Value(t *testing.T) {
	b := []byte(`{"one":1}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	var value int32
	var err error
	if value, err = GetInt32Value(m, "one"); err == nil {
		if value != 1 {
			t.Error("func GetInt32Value return a wrong value (expected for 1)")
		} else {
			t.Log("Get value 1")
		}
	} else {
		t.Error("func GetInt32Value should return no error")
	}
}

func TestGetInt64Value(t *testing.T) {
	b := []byte(`{"one":1}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	var value int64
	var err error
	if value, err = GetInt64Value(m, "one"); err == nil {
		if value != 1 {
			t.Error("func GetInt64Value return a wrong value (expected for 1)")
		} else {
			t.Log("Get value 1")
		}
	} else {
		t.Error("func GetInt64Value should return no error")
	}
}

func TestGetIntValue(t *testing.T) {
	b := []byte(`{"one":1}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	var value int
	var err error
	if value, err = GetIntValue(m, "one"); err == nil {
		if value != 1 {
			t.Error("func GetIntValue return a wrong value (expected for 1)")
		} else {
			t.Log("Get value 1")
		}
	} else {
		t.Error("func GetIntValue should return no error")
	}
}

func TestGetFloat32Value(t *testing.T) {
	b := []byte(`{"pi":3.142}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	var value float32
	var err error
	if value, err = GetFloat32Value(m, "pi"); err == nil {
		if value != 3.142 {
			t.Error("func GetUintValue return a wrong value (expected for 3.142)")
		} else {
			t.Log("Get value 3.142")
		}
	} else {
		t.Error("func GetUintValue should return no error")
	}
}

func TestGetFloat64Value(t *testing.T) {
	b := []byte(`{"pi":3.142}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	var value float64
	var err error
	if value, err = GetFloat64Value(m, "pi"); err == nil {
		if value != 3.142 {
			t.Error("func GetUintValue return a wrong value (expected for 3.142)")
		} else {
			t.Log("Get value 3.142")
		}
	} else {
		t.Error("func GetUintValue should return no error")
	}
}

func TestGetUint8Value(t *testing.T) {
	b := []byte(`{"one":1}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	var value uint8
	var err error
	if value, err = GetUint8Value(m, "one"); err == nil {
		if value != 1 {
			t.Error("func GetUint8Value return a wrong value (expected for 1)")
		} else {
			t.Log("Get value 1")
		}
	} else {
		t.Error(err)
	}
}

func TestGetUint16Value(t *testing.T) {
	b := []byte(`{"one":1}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	var value uint16
	var err error
	if value, err = GetUint16Value(m, "one"); err == nil {
		if value != 1 {
			t.Error("func GetUint16Value return a wrong value (expected for 1)")
		} else {
			t.Log("Get value 1")
		}
	} else {
		t.Error(err)
	}
}

func TestGetUint32Value(t *testing.T) {
	b := []byte(`{"one":1}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	var value uint32
	var err error
	if value, err = GetUint32Value(m, "one"); err == nil {
		if value != 1 {
			t.Error("func GetUint32Value return a wrong value (expected for 1)")
		} else {
			t.Log("Get value 1")
		}
	} else {
		t.Error(err)
	}
}

func TestGetUint64Value(t *testing.T) {
	b := []byte(`{"one":1}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	var value uint64
	var err error
	if value, err = GetUint64Value(m, "one"); err == nil {
		if value != 1 {
			t.Error("func GetUint64Value return a wrong value (expected for 1)")
		} else {
			t.Log("Get value 1")
		}
	} else {
		t.Error(err)
	}
}

func TestGetUintValue(t *testing.T) {
	b := []byte(`{"one":1}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	var value uint
	var err error
	if value, err = GetUintValue(m, "one"); err == nil {
		if value != 1 {
			t.Error("func GetUintValue return a wrong value (expected for 1)")
		} else {
			t.Log("Get value 1")
		}
	} else {
		t.Error(err)
	}
}
