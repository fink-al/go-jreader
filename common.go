package jreader

import "fmt"

// Interface for any type of JSON element
type JSONElement interface {
	// Get element by key (for map) or index (for slice); returns NonExistent if there is no element for the given path
	Get(key any) JSONElement
	// Get any value of element; returns false if there is no value to return for the given path
	Value() (any, bool)
	// Get boolean value of element; returns false if there is no value to return for the given path or if the value is of wrong type
	BooleanValue() (bool, bool)
	// Get number value of element; returns false if there is no value to return for the given path or if the value is of wrong type
	NumberValue() (float64, bool)
	// Get string value of element; returns false if there is no value to return for the given path; otherwise the value present is stringified
	//
	// bool, int and float64 values are stringified using fmt.Sprintf("%v", value)
	// map and slice are stringified using json.Marshal
	StringValue() (string, bool)
}

type JSONData interface {
	[]byte | string | map[string]any | []any | []map[string]any | *[]byte | *string | *map[string]any | *[]any | *[]map[string]any | any
}

// Load JSON data into a JSONElement. Source data may be of type:
//
//	[]byte | string | map[string]any | []any | []map[string]any | *[]byte | *string | *map[string]any | *[]any | *[]map[string]any
func Load[D JSONData](data D) (JSONElement, error) {
	// type switch
	var dany any = data
	switch t := dany.(type) {
	case []byte:
		return fromString(string(t))
	case string:
		return fromString(t)
	case map[string]any:
		return jSONMap(t), nil
	case []any:
		return jSONSlice(t), nil
	case []map[string]any:
		return jSONMapSlice(t), nil
	case *[]byte:
		return fromString(string(safeAccessPointer(t)))
	case *string:
		return fromString(safeAccessPointer(t))
	case *map[string]any:
		return jSONMap(safeAccessPointer(t)), nil
	case *[]any:
		return jSONSlice(safeAccessPointer(t)), nil
	case *[]map[string]any:
		return jSONMapSlice(safeAccessPointer(t)), nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", t)
	}
}

func findTypeOfValue(v any) JSONElement {
	switch v := v.(type) {
	case map[string]any:
		return jSONMap(v)
	case []any:
		return jSONSlice(v)
	case []map[string]any:
		return jSONMapSlice(v)
	case string:
		return jSONValue[string]{value: v}
	case int:
		return jSONValue[int]{value: v}
	case float64:
		return jSONValue[float64]{value: v}
	case bool:
		return jSONValue[bool]{value: v}
	default:
		return nonExistent{}
	}
}
