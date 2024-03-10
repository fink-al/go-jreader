package jreader

import "encoding/json"

func tryParseToSlice(d string) ([]map[string]any, error) {
	var v []map[string]any
	err := json.Unmarshal([]byte(d), &v)
	return v, err
}

func tryParseToMap(d string) (map[string]any, error) {
	var v map[string]any
	err := json.Unmarshal([]byte(d), &v)
	return v, err
}

func fromString(s string) (JSONElement, error) {
	var err error
	var v any
	if v, err = tryParseToMap(s); err == nil {
		return findTypeOfValue(v), nil
	}
	if v, err = tryParseToSlice(s); err == nil {
		return findTypeOfValue(v), nil
	}
	return nil, err
}

func safeAccessPointer[T any](p *T) T {
	if p == nil {
		return getZeroValue[T]()
	}
	return *p
}

func getZeroValue[T any]() T {
	var v T
	return v
}

func toGeneralMap[V any](m map[string]V) map[string]any {
	gm := make(map[string]any)
	for k, v := range m {
		gm[k] = v
	}
	return gm
}
