package mapext

import (
	"errors"
	"reflect"
	"strconv"
)

func convertToString(in interface{}) (string, error) {
	switch val := in.(type) {
	case string:
		return val, nil
	case bool:
		return strconv.FormatBool(val), nil
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64), nil
	case int:
		return strconv.FormatInt(int64(val), 10), nil
	case int8:
		return strconv.FormatInt(int64(val), 10), nil
	case int16:
		return strconv.FormatInt(int64(val), 10), nil
	case int32:
		return strconv.FormatInt(int64(val), 10), nil
	case int64:
		return strconv.FormatInt(val, 10), nil
	case uint:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint64:
		return strconv.FormatUint(val, 10), nil
	}
	return "", errors.New("Unsupported type")
}

func convertFromString(in string, t reflect.Kind) (interface{}, error) {
	switch t {
	case reflect.String:
		return in, nil
	case reflect.Bool:
		return strconv.ParseBool(in)
	case reflect.Float32:
		val, err := strconv.ParseFloat(in, 32)
		return float32(val), err
	case reflect.Float64:
		return strconv.ParseFloat(in, 64)
	case reflect.Int:
		return strconv.Atoi(in)
	case reflect.Int8:
		val, err := strconv.ParseInt(in, 10, 8)
		return int8(val), err
	case reflect.Int16:
		val, err := strconv.ParseInt(in, 10, 16)
		return int16(val), err
	case reflect.Int32:
		val, err := strconv.ParseInt(in, 10, 32)
		return int32(val), err
	case reflect.Int64:
		return strconv.ParseInt(in, 10, 64)
	case reflect.Uint:
		val, err := strconv.ParseUint(in, 10, 0)
		return uint(val), err
	case reflect.Uint8:
		val, err := strconv.ParseUint(in, 10, 8)
		return uint8(val), err
	case reflect.Uint16:
		val, err := strconv.ParseUint(in, 10, 16)
		return uint16(val), err
	case reflect.Uint32:
		val, err := strconv.ParseUint(in, 10, 32)
		return uint32(val), err
	case reflect.Uint64:
		return strconv.ParseUint(in, 10, 64)
	}
	return nil, errors.New("Parse error")
}

func get(m map[string]interface{}, k string, t reflect.Kind) (interface{}, error) {
	v, ok := m[k]
	if !ok {
		return nil, errors.New("Not found the key in map")
	}

	s, err := convertToString(v)
	if err != nil {
		return nil, err
	}

	vi, err := convertFromString(s, t)
	if err != nil {
		return nil, err
	}
	return vi, nil
}

func GetStringValue(m map[string]interface{}, k string) (string, error) {
	var v interface{}
	var err error
	if v, err = get(m, k, reflect.String); nil == err {
		if value, ok := v.(string); ok {
			return value, nil
		}
		return "", errors.New("Assertion error")
	}
	return "", err
}

// GetBoolValue - Search the map for k, and return its value as bool.
// The value in JSON can be
// - bool values: true, True, TRUE, false, False, FALSE
// - int values: 0, 1
// - string values: "true", "True", "TRUE", fa"lse, "False", "FALSE", "0", "1"
func GetBoolValue(m map[string]interface{}, k string) (bool, error) {
	var v interface{}
	var err error
	if v, err = get(m, k, reflect.Bool); nil == err {
		if value, ok := v.(bool); ok {
			return value, nil
		}
		return false, errors.New("Assertion error")
	}
	return false, err
}

// GetIntValue - Search the map for k, and return its value as int.
func GetIntValue(m map[string]interface{}, k string) (int, error) {
	var v interface{}
	var err error
	if v, err = get(m, k, reflect.Int); nil == err {
		if value, ok := v.(int); ok {
			return value, nil
		}
		// 值的类型不匹配，或者超过 int 可表示的范围。
		return 0, errors.New("Assertion error")
	}
	return 0, err
}

// GetInt8Value - Search the map for k, and return its value as int8.
func GetInt8Value(m map[string]interface{}, k string) (int8, error) {
	var v interface{}
	var err error
	if v, err = get(m, k, reflect.Int8); nil == err {
		if value, ok := v.(int8); ok {
			return value, nil
		}
		// 值的类型不匹配，或者超过 int8 可表示的范围。
		return 0, errors.New("Assertion error")
	}
	return 0, err
}

// GetInt16Value - Search the map for k, and return its value as int16.
func GetInt16Value(m map[string]interface{}, k string) (int16, error) {
	var v interface{}
	var err error
	if v, err = get(m, k, reflect.Int16); nil == err {
		if value, ok := v.(int16); ok {
			return value, nil
		}
		// 值的类型不匹配，或者超过 int16 可表示的范围。
		return 0, errors.New("Assertion error")
	}
	return 0, err
}

// GetInt32Value - Search the map for k, and return its value as int32.
func GetInt32Value(m map[string]interface{}, k string) (int32, error) {
	var v interface{}
	var err error
	if v, err = get(m, k, reflect.Int32); nil == err {
		if value, ok := v.(int32); ok {
			return value, nil
		}
		// 值的类型不匹配，或者超过 int32 可表示的范围。
		return 0, errors.New("Assertion error")
	}
	return 0, err
}

// GetInt64Value - Search the map for k, and return its value as int64.
func GetInt64Value(m map[string]interface{}, k string) (int64, error) {
	var v interface{}
	var err error
	if v, err = get(m, k, reflect.Int64); nil == err {
		if value, ok := v.(int64); ok {
			return value, nil
		}
		// 值的类型不匹配，或者超过 int64 可表示的范围。
		return 0, errors.New("Assertion error")
	}
	return 0, err
}

// GetFloat32Value - Search the map for k, and return its value as float32.
func GetFloat32Value(m map[string]interface{}, k string) (float32, error) {
	var v interface{}
	var err error
	if v, err = get(m, k, reflect.Float32); nil == err {
		if value, ok := v.(float32); ok {
			return value, nil
		}
		// 值的类型不匹配，或者超过 float32 可表示的范围。
		return 0, errors.New("Assertion error")
	}
	return 0, err
}

// GetFloat64Value - Search the map for k, and return its value as float64.
func GetFloat64Value(m map[string]interface{}, k string) (float64, error) {
	var v interface{}
	var err error
	if v, err = get(m, k, reflect.Float64); nil == err {
		if value, ok := v.(float64); ok {
			return value, nil
		}
		// 值的类型不匹配，或者超过 float64 可表示的范围。
		return 0, errors.New("Assertion error")
	}
	return 0, err
}

func GetUint8Value(m map[string]interface{}, k string) (uint8, error) {
	var v interface{}
	var err error
	if v, err = get(m, k, reflect.Uint8); nil == err {
		if value, ok := v.(uint8); ok {
			return value, nil
		}
		return 0, errors.New("Assertion error")
	}
	return 0, err
}

func GetUint16Value(m map[string]interface{}, k string) (uint16, error) {
	var v interface{}
	var err error
	if v, err = get(m, k, reflect.Uint16); nil == err {
		if value, ok := v.(uint16); ok {
			return value, nil
		}
		return 0, errors.New("Assertion error")
	}
	return 0, err
}

func GetUint32Value(m map[string]interface{}, k string) (uint32, error) {
	var v interface{}
	var err error
	if v, err = get(m, k, reflect.Uint32); nil == err {
		if value, ok := v.(uint32); ok {
			return value, nil
		}
		return 0, errors.New("Assertion error")
	}
	return 0, err
}

func GetUint64Value(m map[string]interface{}, k string) (uint64, error) {
	var v interface{}
	var err error
	if v, err = get(m, k, reflect.Uint64); nil == err {
		if value, ok := v.(uint64); ok {
			return value, nil
		}
		return 0, errors.New("Assertion error")
	}
	return 0, err
}

func GetUintValue(m map[string]interface{}, k string) (uint, error) {
	var v interface{}
	var err error
	if v, err = get(m, k, reflect.Uint); nil == err {
		if value, ok := v.(uint); ok {
			return value, nil
		}
		return 0, errors.New("Assertion error")
	}
	return 0, err
}
