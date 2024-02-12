package jreader

import "encoding/json"

type jSONMap map[string]any

func (j jSONMap) Get(key any) JSONElement {
	switch key := key.(type) {
	case string:
		if v, ok := j[key]; ok {
			return findTypeOfValue(v)
		}
	default:
		return nonExistent{}
	}
	return nonExistent{}
}

func (j jSONMap) Value() (any, bool) {
	return j, true
}

func (j jSONMap) BooleanValue() (bool, bool) {
	return false, false
}

func (j jSONMap) NumberValue() (float64, bool) {
	return 0, false
}

func (j jSONMap) StringValue() (string, bool) {
	jsonBytes, err := json.Marshal(j)
	if err != nil {
		return "", false
	}
	return string(jsonBytes), true
}
