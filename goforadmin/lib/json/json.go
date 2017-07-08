package json

import "encoding/json"

func JsonEncode(v interface{}) string {
	bytes, _ := json.Marshal(v)

	return string(bytes)
}
