package utils

import "encoding/json"

func UnmarshalUint64(message json.RawMessage) (uint64, error) {
	var tmp uint64

	if err := json.Unmarshal(message, &tmp); err != nil {
		return 0, err
	}

	return tmp, nil
}

func UnmarshalString(message json.RawMessage) (string, error) {
	var tmp string

	if err := json.Unmarshal(message, &tmp); err != nil {
		return "", err
	}

	return tmp, nil
}
