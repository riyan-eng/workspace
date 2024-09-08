package util

import "net/url"

type convert struct{}

func NewConvert() *convert {
	return &convert{}
}

func (m *convert) AnyToStr(a any) string {
	var b = ""
	switch a := a.(type) {
	case url.Values:
		b = a.Encode()
	case string:
		b = a
	}

	return b
}
