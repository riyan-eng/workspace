package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertAnyToString(t *testing.T) {
	expect := "zombie"
	actual := NewConvert().AnyToStr("zombie")
	assert.Equal(t,
		expect,
		actual,
	)
}
