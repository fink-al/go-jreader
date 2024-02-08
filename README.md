# go-jreader

[![goversion-tag-do-not-edit](https://img.shields.io/badge/Go%20Version-1.18-blue.svg)](https://shields.io/)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-49%25-brightgreen.svg?longCache=true&style=flat)</a>

Safely read specific values from nested JSON objects, maps or slices.
Small, zero dependency wrapper interface that allows to use optional chaining in go.

## Usage

```golang
jsonObj, _ := Load("{\"lvl1_nested_obj\": {\"lvl2_nested_list\": [{\"key_of_value\": \"value\"}]}}")
va, ok := jsonObj.Get("lvl1_nested_obj").Get("lvl2_nested_list").Get(0).Get("key_of_value").StringValue()
```
