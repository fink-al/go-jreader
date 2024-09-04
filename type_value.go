package jreader

import (
	"fmt"
	"strconv"
)

type jSONValue[T any] struct {
	value T
}

func (j jSONValue[T]) Get(key any) JSONElement { //nolint: revive // interface implementation method
	return nonExistent{}
}

func (j jSONValue[T]) Value() (any, bool) {
	return j.value, true
}

func (j jSONValue[T]) BooleanValue() (bool, bool) {
	var vi any = j.value
	if v, ok := vi.(bool); ok {
		return v, true
	}
	return false, false
}

func (j jSONValue[T]) NumberValue() (float64, bool) {
	var vi any = j.value
	// convert any number value to float64
	var y float64
	found := true
	switch v := vi.(type) {
	case int:
		y = float64(v)
	case int8:
		y = float64(v)
	case int16:
		y = float64(v)
	case int32:
		y = float64(v)
	case int64:
		y = float64(v)
	case uint:
		y = float64(v)
	case uint8:
		y = float64(v)
	case uint16:
		y = float64(v)
	case uint32:
		y = float64(v)
	case uint64:
		y = float64(v)
	case float32:
		y = float64(v)
	case float64:
		y = v
	case string:
		// try convert string to float64
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			found = false
		} else {
			y = f
		}
	default:
		found = false
	}
	return y, found
}

func (j jSONValue[T]) StringValue() (string, bool) {
	var vi any = j.value
	if v, ok := vi.(string); ok {
		return v, true
	} else if v, ok := vi.(bool); ok {
		return fmt.Sprintf("%v", v), true
	} else if v, ok := vi.(float64); ok {
		return fmt.Sprintf("%v", v), true
	}
	return "", false
}

func (j jSONValue[T]) MapValue() (map[string]any, bool) {
	return map[string]any{}, false
}

func (j jSONValue[T]) SliceValue() ([]any, bool) {
	return []any{}, false
}

func (n jSONValue[T]) MapJSONElementValue() (map[string]JSONElement, bool) {
	mvn := map[string]JSONElement{}
	return mvn, false
}

func (n jSONValue[T]) SliceJSONElementValue() ([]JSONElement, bool) {
	return []JSONElement{}, false
}
