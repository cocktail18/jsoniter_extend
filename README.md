
An extention for `github.com/json-iterator/go`

# Usage

## treat {} as empty array
```go
var json = jsoniter.ConfigCompatibleWithStandardLibrary
//  treat {} as empty array
jsoniter.RegisterExtension(&TolerateEmptyStructExtension{})
json.Marshal(&data)
```