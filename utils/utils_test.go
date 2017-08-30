package utils

import (
	"reflect"
	"testing"
)

func TestArrayOfBytes(t *testing.T) {
	var p []byte
	for i := 3; i != 0; i-- {
		p = append(p, 0)
	}
	if r := ArrayOfBytes(3, 0); !reflect.DeepEqual(r, p) {
		t.Fatalf("ArrayOfBytes(3, 0) expect [0 0 0], but %v", p)
	}
}
