package jreader

type (
	JSONSlice    []any
	JSONMapSlice []map[string]any
)

func (j JSONSlice) Get(key any) JSONElement {
	switch key := key.(type) {
	case int:
		if key < len(j) {
			return findTypeOfValue(j[key])
		} else {
			return NonExistent{}
		}
	default:
		return NonExistent{}
	}
}

func (j JSONSlice) Value() (any, bool) {
	return j, true
}

func (j JSONMapSlice) BooleanValue() (bool, bool) {
	return false, false
}

func (j JSONMapSlice) NumberValue() (float64, bool) {
	return 0, false
}

func (j JSONMapSlice) StringValue() (string, bool) {
	return "", false
}

func (j JSONMapSlice) Get(key any) JSONElement {
	switch key := key.(type) {
	case int:
		if key < len(j) {
			return findTypeOfValue(j[key])
		} else {
			return NonExistent{}
		}
	default:
		return NonExistent{}
	}
}

func (j JSONMapSlice) Value() (any, bool) {
	return j, true
}

func (j JSONSlice) BooleanValue() (bool, bool) {
	return false, false
}

func (j JSONSlice) NumberValue() (float64, bool) {
	return 0, false
}

func (j JSONSlice) StringValue() (string, bool) {
	return "", false
}
