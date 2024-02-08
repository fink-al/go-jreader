package jreader

type JSONValue[T any] struct {
	value T
}

func (j JSONValue[T]) Get(key any) JSONElement {
	return NonExistent{}
}

func (j JSONValue[T]) Value() (any, bool) {
	return j.value, true
}

func (j JSONValue[T]) BooleanValue() (bool, bool) {
	var vi any = j.value
	if v, ok := vi.(bool); ok {
		return v, true
	}
	return false, false
}

func (j JSONValue[T]) NumberValue() (float64, bool) {
	var vi any = j.value
	if v, ok := vi.(float64); ok {
		return v, true
	}
	return 0, false
}

func (j JSONValue[T]) StringValue() (string, bool) {
	var vi any = j.value
	if v, ok := vi.(string); ok {
		return v, true
	}
	return "", false
}
