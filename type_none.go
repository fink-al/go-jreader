package jreader

type (
	NonExistent struct{}
)

func (n NonExistent) Get(key any) JSONElement {
	return NonExistent{}
}

func (n NonExistent) Value() (any, bool) {
	return nil, false
}

func (n NonExistent) BooleanValue() (bool, bool) {
	return false, false
}

func (n NonExistent) NumberValue() (float64, bool) {
	return 0, false
}

func (n NonExistent) StringValue() (string, bool) {
	return "", false
}
