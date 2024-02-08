package jreader

type JSONMap map[string]any

func (j JSONMap) Get(key any) JSONElement {
	switch key := key.(type) {
	case string:
		if v, ok := j[key]; ok {
			return findTypeOfValue(v)
		}
	default:
		return NonExistent{}
	}
	return NonExistent{}
}

func (j JSONMap) Value() (any, bool) {
	return j, true
}

func (j JSONMap) BooleanValue() (bool, bool) {
	return false, false
}

func (j JSONMap) NumberValue() (float64, bool) {
	return 0, false
}

func (j JSONMap) StringValue() (string, bool) {
	return "", false
}
