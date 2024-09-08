package util

import "github.com/bytedance/sonic"

func UnmarshalConverter[T any](s string) (data T) {
	sonic.UnmarshalString(s, &data)
	return
}
