package util

import (
	"testing"
)

func TestGetSign(t *testing.T) {
	sign := GetSign(map[string]interface{}{}, "c")
	t.Log(sign)
}
