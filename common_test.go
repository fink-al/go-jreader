package jreader

import (
	"fmt"
	"testing"
)

func TestLoadString(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name      string
		args      args
		checkfunc func(t *testing.T, got JSONElement, gotErr error) error
		wantErr   bool
	}{
		{"1: single layer hierarchy, single element int", args{"{\"a\": 1}"}, func(t *testing.T, got JSONElement, gotErr error) error {
			if v, ok := got.Get("a").Value(); !ok || v != float64(1) {
				t.Errorf("Load() = %v, want %v", v, 1)
				return fmt.Errorf("Load() = %v, want %v", v, 1)
			}
			return nil
		}, false},
		{"2: two layer hierarchy, single element int", args{"{\"a\": {\"b\": 1}}"}, func(t *testing.T, got JSONElement, gotErr error) error {
			if v, ok := got.Get("a").Get("b").NumberValue(); !ok || v != float64(1) {
				t.Errorf("Load() = %v, want %v", v, 1)
				return fmt.Errorf("Load() = %v, want %v", v, 1)
			}
			return nil
		}, false},
		{"3: single layer hierarchy, single element with list value int", args{"{\"a\": [1, 2, 3]}"}, func(t *testing.T, got JSONElement, gotErr error) error {
			if v, ok := got.Get("a").Get(1).Value(); !ok || v != float64(2) {
				t.Errorf("Load() = %v %T, want %v %T", v, v, 2, float64(2))
				return fmt.Errorf("Load() = %v, want %v", v, 2)
			}
			return nil
		}, false},
		{"4: single layer hierarchy, single element with list value int", args{"{\"a\": [1, 2, 3]}"}, func(t *testing.T, got JSONElement, gotErr error) error {
			if v, ok := got.Get("b").Get(1).Value(); ok || v != nil {
				t.Errorf("Load() = %v, want %v", v, nil)
				return fmt.Errorf("Load() = %v, want %v", v, nil)
			}
			return nil
		}, false},
		{"5: string value", args{"{\"a\": \"b\"}"}, func(t *testing.T, got JSONElement, gotErr error) error {
			if v, ok := got.Get("a").StringValue(); !ok || v != "b" {
				t.Errorf("Load() = %v, want %v", v, "b")
				return fmt.Errorf("Load() = %v, want %v", v, "b")
			}
			return nil
		}, false},
		{"6: boolean value", args{"{\"a\": true}"}, func(t *testing.T, got JSONElement, gotErr error) error {
			if v, ok := got.Get("a").BooleanValue(); !ok || v != true {
				t.Errorf("Load() = %v, want %v", v, true)
				return fmt.Errorf("Load() = %v, want %v", v, true)
			}
			return nil
		}, false},
		{"7: boolean value", args{"{\"a\": false}"}, func(t *testing.T, got JSONElement, gotErr error) error {
			if v, ok := got.Get("a").BooleanValue(); !ok || v != false {
				t.Errorf("Load() = %v, want %v", v, false)
				return fmt.Errorf("Load() = %v, want %v", v, false)
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
		`}, func(t *testing.T, got JSONElement, gotErr error) error {
			if v, ok := got.Get("glossary").Get("GlossDiv").Get("GlossList").Get("GlossEntry").Get("GlossDef").Get("GlossSeeAlso").Get(1).StringValue(); !ok || v != "XML" {
				t.Errorf("Load() = %v, want %v", v, "XML")
				return fmt.Errorf("Load() = %v, want %v", v, "XML")
			}
			// test entry that does not exist
			if v, ok := got.Get("glossary").Get("GlossDiv").Get("GlossList").Get("GlossEntry").Get("GlossDef").Get("GlossSeeAlso").Get(2).StringValue(); ok || v != "" {
				t.Errorf("Load() = %v, want %v", v, "")
				return fmt.Errorf("Load() = %v, want %v", v, "")
			}
			// test entry that does not exis
			if v, ok := got.Get("glossary").Get("GlossDivWRONG").Get("GlossList").Get("GlossEntry").Get("GlossDef").Get("GlossSeeAlso").Get(2).StringValue(); ok || v != "" {
				t.Errorf("Load() = %v, want %v", v, "")
				return fmt.Errorf("Load() = %v, want %v", v, "")
			}
			return nil
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, er := Load(tt.args.data)
			if (er != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", er, tt.wantErr)
				return
			}
			if err := tt.checkfunc(t, got, er); err != nil {
				t.Error(err)
			}
		})
	}
}
