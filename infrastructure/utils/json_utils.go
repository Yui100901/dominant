package utils

import "encoding/json"

// @Author yfy2001
// @Date 2024/9/13 16 51

func Unmarshal[T any](data []byte) (*T, error) {
	s := new(T)
	err := json.Unmarshal(data, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
