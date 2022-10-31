package util

import (
	"testing"
)

func TestGetSign(t *testing.T) {
	sign := GetSign(map[string]interface{}{"timestamp": 1}, map[string]interface{}{"a": "b"}, "c")
	t.Log(sign)
}
