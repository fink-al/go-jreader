# go-jreader

[![goversion-tag-do-not-edit](https://img.shields.io/badge/Go%20Version-1.18-blue.svg)](https://shields.io/)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-65%25-brightgreen.svg?longCache=true&style=flat)</a>

Safely read specific values from nested JSON objects, maps or slices.
Small, zero dependency wrapper interface that allows to use optional chaining in go.

## Prerequisites

Requires go1.18 (generics)

## Install

```bash
go get github.com/fink-al/go-jreader
```

## Usage Examples

```golang
// Load JSON data into a JSONElement. Source data may be of type:
//
//  []byte | string | map[string]any | []any |
//  []map[string]any | *[]byte | *string |
//  *map[string]any | *[]any | *[]map[string]any
//
data, err := jreader.Load("{\"a\": 1}")
data.Get("a").NumberValue() // returns 1, true
data.Get("b").NumberValue() // returns 0, false
data.Get("a").StringValue() // returns "1", true

data, err = jreader.Load([]byte("[{\"a\": 1}]"))
data.Get(0).Get("a").NumberValue() // returns 1, true

data, err = jreader.Load(map[string]any{"a": 1})
data.Get("a").NumberValue() // returns 1, true

data, err = jreader.Load([]any{1, 2, 3})
data.Get(1).NumberValue() // returns 2, true

data, err = jreader.Load([]map[string]any{{"a": 1}, {"b": 2}})
data.Get(1).Get("b").NumberValue() // returns 2, true
data.Get(1).StringValue()          // returns "{\"b\":2}", true

data, err = jreader.Load(createPointer([]byte("{\"nest_lvl1\": {\"nest_lvl2\": 1}}")))
data.Get("nest_lvl1").Get("nest_lvl2").NumberValue() // returns 1, true
data.Get("INVALID").Get("nest_lvl2").NumberValue()   // returns 0, false

jsonObj, _ := jreader.Load("{\"lvl1_nested_obj\": {\"lvl2_nested_list\": [{\"key_of_value\": \"value\"}]}}")
va, ok := jsonObj.Get("lvl1_nested_obj").
  Get("lvl2_nested_list").Get(0).
  Get("key_of_value").StringValue() // returns "value", true

```
