package jreader

import (
	"encoding/json"
	"log"
)

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
		log.Default().Printf("Error marshalling JSONMap: %s", err.Error())
		return "", false
	}
	return string(jsonBytes), true
}

func (j jSONMap) MapValue() (map[string]any, bool) {
	return j, true
}

func (j jSONMap) SliceValue() ([]any, bool) {
	return []any{}, false
}
