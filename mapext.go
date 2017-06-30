package mapext

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
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
		return strconv.ParseUint(in, 10, 0)
	case reflect.Uint8:
		return strconv.ParseUint(in, 10, 8)
	case reflect.Uint16:
		return strconv.ParseUint(in, 10, 16)
	case reflect.Uint32:
		return strconv.ParseUint(in, 10, 32)
	case reflect.Uint64:
		return strconv.ParseUint(in, 10, 64)
	}
	return nil, errors.New("Parse error")
}

func get(m map[string]interface{}, k string, t reflect.Kind) (interface{}, error) {
	for key, val := range m {
		if strings.Compare(k, key) == 0 {
			var err error
			if valInString, err := convertToString(val); nil == err {
				if value, err := convertFromString(valInString, t); nil == err {
					return value, nil
				}
			}
			return nil, err
		}
	}
	return nil, errors.New("Not found the key in map")
}

func GetStringValue(m map[string]interface{}, k string) (string, error) {
	var err error
	if v, err := get(m, k, reflect.String); nil == err {
		if value, ok := v.(string); ok {
			return value, nil
		}
		return "", errors.New("Assertion error")
	}
	return "", err
}

func GetIntValue(m map[string]interface{}, k string) (int, error) {
	var err error
	if v, err := get(m, k, reflect.Int); nil == err {
		if value, ok := v.(int); ok {
			return value, nil
		}
		// 值的类型不匹配，或者超过 int 可表示的范围。
		return 0, errors.New("Assertion error")
	}
	return 0, err
}

func GetInt8Value(m map[string]interface{}, k string) (int8, error) {
	var err error
	if v, err := get(m, k, reflect.Int8); nil == err {
		if value, ok := v.(int8); ok {
			return value, nil
		}
		// 值的类型不匹配，或者超过 int8 可表示的范围。
		return 0, errors.New("Assertion error")
	}
	return 0, err
}

func GetInt16Value(m map[string]interface{}, k string) (int16, error) {
	var err error
	if v, err := get(m, k, reflect.Int16); nil == err {
		if value, ok := v.(int16); ok {
			return value, nil
		}
		// 值的类型不匹配，或者超过 int16 可表示的范围。
		return 0, errors.New("Assertion error")
	}
	return 0, err
}

func GetInt32Value(m map[string]interface{}, k string) (int32, error) {
	var err error
	if v, err := get(m, k, reflect.Int32); nil == err {
		if value, ok := v.(int32); ok {
			return value, nil
		}
		// 值的类型不匹配，或者超过 int32 可表示的范围。
		return 0, errors.New("Assertion error")
	}
	return 0, err
}

func GetInt64Value(m map[string]interface{}, k string) (int64, error) {
	var err error
	if v, err := get(m, k, reflect.Int64); nil == err {
		if value, ok := v.(int64); ok {
			return value, nil
		}
		// 值的类型不匹配，或者超过 int64 可表示的范围。
		return 0, errors.New("Assertion error")
	}
	return 0, err
}
