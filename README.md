# Encoding Json
Golang standard `encoding/json` module with more options.

## Options
- `in`: used to include a field only in decoding
- `out`: used to include a field only in encoding

## Description
| example            | description                               |
|--------------------|-------------------------------------------|
| json:"-"           | field is ignored                          |
| json:"name"        | field is included in `input` and `output` |
| json:"name,in"     | field is only included in `input`         |
| json:"name,out"    | field is only included in `output`        |
| json:"name,in,out" | field is included in `input` and `output` | 


## Example
Considering the following struct as an example:
- included in decode(unmarshalling):
  - name
  - email
  - password
- included in encode(marshalling):
  - name
  - email
  - balance

```go
package main

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,in"`
	Balance  uint   `json:"balance,out"`
}
```

# Acknowledgment
This module is a copy of `encoding/json` module from Go version `1.19`.
