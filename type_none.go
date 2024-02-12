package jreader

type (
	nonExistent struct{}
)

func (n nonExistent) Get(key any) JSONElement { //nolint: revive // interface implementation method
	return nonExistent{}
}

func (n nonExistent) Value() (any, bool) {
	return nil, false
}

func (n nonExistent) BooleanValue() (bool, bool) {
	return false, false
}

func (n nonExistent) NumberValue() (float64, bool) {
	return 0, false
}

func (n nonExistent) StringValue() (string, bool) {
	return "", false
}
