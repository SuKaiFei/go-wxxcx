package util

import (
	"testing"
)

func TestGetSign(t *testing.T) {
	sign := GetSign(map[string][]string{}, "c")
	t.Log(sign)
}
