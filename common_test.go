package jreader_test

import (
	"fmt"
	"testing"

	"github.com/fink-al/go-jreader"
)

func TestLoadString(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name      string
		args      args
		checkfunc func(t *testing.T, got jreader.JSONElement, gotErr error) error
		wantErr   bool
	}{
		{"1: single layer hierarchy, single element int", args{"{\"a\": 1}"}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").Value(); !ok || v != float64(1) {
				t.Errorf("jreader.Load() = %v, want %v", v, 1)
				return fmt.Errorf("jreader.Load() = %v, want %v", v, 1)
			}
			return nil
		}, false},
		{"2: two layer hierarchy, single element int", args{"{\"a\": {\"b\": 1}}"}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").Get("b").NumberValue(); !ok || v != float64(1) {
				t.Errorf("jreader.Load() = %v, want %v", v, 1)
				return fmt.Errorf("jreader.Load() = %v, want %v", v, 1)
			}
			return nil
		}, false},
		{"3: single layer hierarchy, single element with list value int", args{"{\"a\": [1, 2, 3]}"}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").Get(1).Value(); !ok || v != float64(2) {
				t.Errorf("jreader.Load() = %v %T, want %v %T", v, v, 2, float64(2))
				return fmt.Errorf("jreader.Load() = %v, want %v", v, 2)
			}
			return nil
		}, false},
		{"3: single layer hierarchy, single element with list value int", args{"{\"a\": [1, 2, 3]}"}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").SliceValue(); !ok || len(v) != 3 {
				t.Errorf("jreader.Load() = %v %T, want %v %T", v, v, 3, 3)
				return fmt.Errorf("jreader.Load() = %v, want %v", v, 2)
			}
			return nil
		}, false},
		{"4: single layer hierarchy, single element with list value int", args{"{\"a\": [1, 2, 3]}"}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("b").Get(1).Value(); ok || v != nil {
				t.Errorf("jreader.Load() = %v, want %v", v, nil)
				return fmt.Errorf("jreader.Load() = %v, want %v", v, nil)
			}
			return nil
		}, false},
		{"5: string value", args{"{\"a\": \"b\"}"}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").StringValue(); !ok || v != "b" {
				t.Errorf("jreader.Load() = %v, want %v", v, "b")
				return fmt.Errorf("jreader.Load() = %v, want %v", v, "b")
			}
			return nil
		}, false},
		{"6: boolean value", args{"{\"a\": true}"}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").BooleanValue(); !ok || v != true {
				t.Errorf("jreader.Load() = %v, want %v", v, true)
				return fmt.Errorf("jreader.Load() = %v, want %v", v, true)
			}
			return nil
		}, false},
		{"7: boolean value", args{"{\"a\": false}"}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").BooleanValue(); !ok || v != false {
				t.Errorf("jreader.Load() = %v, want %v", v, false)
				return fmt.Errorf("jreader.Load() = %v, want %v", v, false)
			}
			return nil
		}, false},
		{"8: boolean value", args{`
		{
			"glossary": {
					"title": "example glossary",
			"GlossDiv": {
							"title": "S",
				"GlossList": {
									"GlossEntry": {
											"ID": "SGML",
						"SortAs": "SGML",
						"GlossTerm": "Standard Generalized Markup Language",
						"Acronym": "SGML",
						"Abbrev": "ISO 8879:1986",
						"GlossDef": {
													"para": "A meta-markup language, used to create markup languages such as DocBook.",
							"GlossSeeAlso": ["GML", "XML"]
											},
						"GlossSee": "markup"
									}
							}
					}
			}
	}
		`}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("glossary").Get("GlossDiv").Get("GlossList").Get("GlossEntry").Get("GlossDef").Get("GlossSeeAlso").Get(1).StringValue(); !ok || v != "XML" {
				t.Errorf("jreader.Load() = %v, want %v", v, "XML")
				return fmt.Errorf("jreader.Load() = %v, want %v", v, "XML")
			}
			// test entry that does not exist
			if v, ok := got.Get("glossary").Get("GlossDiv").Get("GlossList").Get("GlossEntry").Get("GlossDef").Get("GlossSeeAlso").Get(2).StringValue(); ok || v != "" {
				t.Errorf("jreader.Load() = %v, want %v", v, "")
				return fmt.Errorf("jreader.Load() = %v, want %v", v, "")
			}
			// test entry that does not exis
			if v, ok := got.Get("glossary").Get("GlossDivWRONG").Get("GlossList").Get("GlossEntry").Get("GlossDef").Get("GlossSeeAlso").Get(2).StringValue(); ok || v != "" {
				t.Errorf("jreader.Load() = %v, want %v", v, "")
				return fmt.Errorf("jreader.Load() = %v, want %v", v, "")
			}
			return nil
		}, false},
		{"9: boolean value", args{"{\"a\": null}"}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").StringValue(); ok || v != "" {
				t.Errorf("jreader.Load() = %v, want %v", v, "")
				return fmt.Errorf("jreader.Load() = %v, want %v", v, "")
			}
			return nil
		}, false},
		{"10: jreader.Load bool value as string", args{"{\"a\": true}"}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").StringValue(); !ok || v != "true" {
				t.Errorf("jreader.Load() = %v, want %v", v, "")
				return fmt.Errorf("jreader.Load() = %v, want %v", v, "")
			}
			return nil
		}, false},
		{"11: jreader.Load int value as string", args{"{\"a\": 1}"}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").StringValue(); !ok || v != "1" {
				t.Errorf("jreader.Load() = %v, want %v", v, "")
				return fmt.Errorf("jreader.Load() = %v, want %v", v, "")
			}
			return nil
		}, false},
		{"12: jreader.Load float value as string", args{"{\"a\": 1.1}"}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").StringValue(); !ok || v != "1.1" {
				t.Errorf("jreader.Load() = %v, want %v", v, "")
				return fmt.Errorf("jreader.Load() = %v, want %v", v, "")
			}
			return nil
		}, false},
		{"13: jreader.Load map value as string", args{"{\"a\": {\"b\": 1}}"}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").StringValue(); !ok || v != "{\"b\":1}" {
				t.Errorf("jreader.Load() = %v, want %v", v, "")
				return fmt.Errorf("jreader.Load() = %v, want %v", v, "")
			}
			return nil
		}, false},
		{"14: jreader.Load list value as string", args{"{\"a\": [1, 2, 3] }"}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").StringValue(); !ok || v != "[1,2,3]" {
				t.Errorf("jreader.Load() = %v, want %v", v, "")
				return fmt.Errorf("jreader.Load() = %v, want %v", v, "")
			}
			return nil
		}, false},
		{"15: jreader.Load list of map value as string", args{"{\"a\": [{\"b\": 1}, {\"c\": 2}] }"}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").StringValue(); !ok || v != "[{\"b\":1},{\"c\":2}]" {
				t.Errorf("jreader.Load() = %v, want %v", v, "")
				return fmt.Errorf("jreader.Load() = %v, want %v", v, "")
			}
			return nil
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, er := jreader.Load(tt.args.data)
			if (er != nil) != tt.wantErr {
				t.Errorf("jreader.Load() error = %v, wantErr %v", er, tt.wantErr)
				return
			}
			if err := tt.checkfunc(t, got, er); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestLoadDifferentTypes(t *testing.T) {
	type args struct {
		data any
	}
	tests := []struct {
		name      string
		args      args
		checkfunc func(t *testing.T, got jreader.JSONElement, gotErr error) error
		wantErr   bool
	}{
		{"1: jreader.Load string", args{"{\"a\": 1}"}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").NumberValue(); !ok || v != float64(1) {
				t.Errorf("jreader.Load() = %v, want %v", v, 1)
				return fmt.Errorf("jreader.Load() = %v, want %v", v, 1)
			}
			return nil
		}, false},
		{"2: jreader.Load string", args{[]byte("{\"a\": 1}")}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").NumberValue(); !ok || v != float64(1) {
				t.Errorf("jreader.Load() = %v, want %v", v, 1)
				return fmt.Errorf("jreader.Load() = %v, want %v", v, 1)
			}
			return nil
		}, false},
		{"3: jreader.Load string", args{map[string]any{"a": 1}}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").NumberValue(); !ok || v != float64(1) {
				t.Errorf("jreader.Load() = %v, want %v", v, 1)
				return fmt.Errorf("jreader.Load() = %v, want %v", v, 1)
			}
			return nil
		}, false},
		{"4: jreader.Load string", args{[]any{1, 2, 3}}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get(1).NumberValue(); !ok || v != float64(2) {
				t.Errorf("jreader.Load() = %v, want %v", v, 2)
				return fmt.Errorf("jreader.Load() = %v, want %v", v, 2)
			}
			return nil
		}, false},
		{"5: jreader.Load string", args{[]map[string]any{{"a": 1}, {"b": 2}}}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get(1).Get("b").NumberValue(); !ok || v != float64(2) {
				t.Errorf("jreader.Load() = %v, want %v", v, 2)
				return fmt.Errorf("jreader.Load() = %v, want %v", v, 2)
			}
			return nil
		}, false},
		{"6: jreader.Load string", args{createPointer([]byte("{\"a\": 1}"))}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").NumberValue(); !ok || v != float64(1) {
				t.Errorf("jreader.Load() = %v, want %v", v, 1)
				return fmt.Errorf("jreader.Load() = %v, want %v", v, 1)
			}
			return nil
		}, false},
		{"7: jreader.Load invalid string", args{createPointer("\"a\": 1")}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if got != nil || gotErr == nil {
				t.Errorf("jreader.Load() = %v, want %v", got, nil)
				return fmt.Errorf("jreader.Load() = %v, want %v", got, nil)
			}
			return nil
		}, true},
		{"8: jreader.Load string", args{createPointer(map[string]any{"a": 1})}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").NumberValue(); !ok || v != float64(1) {
				t.Errorf("jreader.Load() = %v, want %v", v, 1)
				return fmt.Errorf("jreader.Load() = %v, want %v", v, 1)
			}
			return nil
		}, false},
		{"9: jreader.Load string", args{&[]any{1, 2, 3}}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get(1).NumberValue(); !ok || v != float64(2) {
				t.Errorf("jreader.Load() = %v, want %v", v, 2)
				return fmt.Errorf("jreader.Load() = %v, want %v", v, 2)
			}
			return nil
		}, false},
		{"10: jreader.Load string", args{&[]map[string]any{{"a": 1}, {"b": 2}}}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get(1).Get("b").NumberValue(); !ok || v != float64(2) {
				t.Errorf("jreader.Load() = %v, want %v", v, 2)
				return fmt.Errorf("jreader.Load() = %v, want %v", v, 2)
			}
			return nil
		}, false},
		{"10: jreader.Load slice", args{&[]map[string]any{{"a": 1}, {"b": 2}}}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.SliceValue(); !ok || len(v) != 2 {
				t.Errorf("jreader.Load() = %v, want %v", v, 2)
				return fmt.Errorf("jreader.Load() = %v, want %v", v, 2)
			}
			return nil
		}, false},
		{"11: jreader.Load string", args{map[string]string{"a": "1"}}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.Get("a").NumberValue(); !ok || v != float64(1) {
				t.Errorf("jreader.Load() = %v, want %v", v, 1)
				return fmt.Errorf("jreader.Load() = %v, want %v", v, 1)
			}
			return nil
		}, false},
		{"11: jreader.Load map", args{map[string]string{"a": "1"}}, func(t *testing.T, got jreader.JSONElement, gotErr error) error {
			if v, ok := got.MapValue(); !ok || v["a"] != "1" {
				t.Errorf("jreader.Load() = %v, want %v", v, 1)
				return fmt.Errorf("jreader.Load() = %v, want %v", v, 1)
			}
			return nil
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, er := jreader.Load(tt.args.data)
			if (er != nil) != tt.wantErr {
				t.Errorf("jreader.Load() error = %v, wantErr %v", er, tt.wantErr)
				return
			}
			if err := tt.checkfunc(t, got, er); err != nil {
				t.Error(err)
			}
		})
	}
}

func createPointer[T any](v T) *T {
	return &v
}
