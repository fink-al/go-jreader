package jreader

import "encoding/json"

type (
	jSONSlice    []any
	jSONMapSlice []map[string]any
)

func (j jSONSlice) Get(key any) JSONElement {
	switch key := key.(type) {
	case int:
		if key < len(j) {
			return findTypeOfValue(j[key])
		}
		return nonExistent{}
	default:
		return nonExistent{}
	}
}

func (j jSONSlice) Value() (any, bool) {
	return j, true
}

func (j jSONMapSlice) BooleanValue() (bool, bool) {
	return false, false
}

func (j jSONMapSlice) NumberValue() (float64, bool) {
	return 0, false
}

func (j jSONMapSlice) StringValue() (string, bool) {
	return "", false
}

func (j jSONMapSlice) Get(key any) JSONElement {
	switch key := key.(type) {
	case int:
		if key < len(j) {
			return findTypeOfValue(j[key])
		}
		return nonExistent{}
	default:
		return nonExistent{}
	}
}

func (j jSONMapSlice) Value() (any, bool) {
	return j, true
}

func (j jSONSlice) BooleanValue() (bool, bool) {
	return false, false
}

func (j jSONSlice) NumberValue() (float64, bool) {
	return 0, false
}

func (j jSONSlice) StringValue() (string, bool) {
	jsonBytes, err := json.Marshal(j)
	if err != nil {
		return "", false
	}
	return string(jsonBytes), true
}
