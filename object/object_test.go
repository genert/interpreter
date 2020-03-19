package object_test

import (
	"genert.org/interpreter/object"
	"testing"
)

func TestStringHashKey(t *testing.T) {
	hello1 := &object.String{Value: "Hello World"}
	hello2 := &object.String{Value: "Hello World"}
	diff1 := &object.String{Value: "My name is johnny"}
	diff2 := &object.String{Value: "My name is johnny"}
	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have")
	}
	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have")
	}
	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content")
	}
}
