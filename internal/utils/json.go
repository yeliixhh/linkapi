package utils

import "encoding/json"

func ToJSON(v interface{}) string {
	marshal, err := json.Marshal(v)

	if err != nil {
		return ""
	}

	return string(marshal)
}
