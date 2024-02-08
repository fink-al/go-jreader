package jreader

import "fmt"

type JSONElement interface {
	// Get element by key (string or int)
	Get(key any) JSONElement
	// Get any value of element; returns false if there is no value to return for the given path
	Value() (any, bool)
	// Get boolean value of element; returns false if there is no value to return for the given path or if the value is of wrong type
	BooleanValue() (bool, bool)
	// Get number value of element; returns false if there is no value to return for the given path or if the value is of wrong type
	NumberValue() (float64, bool)
	// Get string value of element;  returns false if there is no value to return for the given path or if the value is of wrong type
	StringValue() (string, bool)
}

type JSONData interface {
	byte | string | map[string]any | []any | []map[string]any
}

// Load JSON data into a JSONElement.
// Source data may be of type string, byte slice, map or slice.
func Load[D JSONData](data D) (JSONElement, error) {
	// type switch
	var dany any = data
	switch t := dany.(type) {
	case []byte:
		return fromString(string(t))
	case string:
		return fromString(t)
	case map[string]any:
		return JSONMap(t), nil
	case []any:
		return JSONSlice(t), nil
	case []map[string]any:
		return JSONMapSlice(t), nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", t)
	}
}

func findTypeOfValue(v any) JSONElement {
	switch v := v.(type) {
	case map[string]any:
		return JSONMap(v)
	case []any:
		return JSONSlice(v)
	case []map[string]any:
		return JSONMapSlice(v)
	case string:
		return JSONValue[string]{value: v}
	case int:
		return JSONValue[int]{value: v}
	case float64:
		return JSONValue[float64]{value: v}
	case bool:
		return JSONValue[bool]{value: v}
	default:
		return NonExistent{}
	}
}
