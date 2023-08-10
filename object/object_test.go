// object/object_test.go

package object

import (
	"testing"
)

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	true1 := &Boolean{Value: true}
	true2 := &Boolean{Value: true}
	false1 := &Boolean{Value: false}
	false2 := &Boolean{Value: false}

	if true1.HashKey() != true2.HashKey() {
		t.Errorf("booleans with same content have different hash keys")
	}

	if false1.HashKey() != false2.HashKey() {
		t.Errorf("booleans with same content have different hash keys")
	}

	if true1.HashKey() == false1.HashKey() {
		t.Errorf("booleans with different content have same hash keys")
	}
}

func TestIntegerHashKey(t *testing.T) {
	int1 := &Integer{Value: 1}
	int2 := &Integer{Value: 1}
	int3 := &Integer{Value: 2}

	if int1.HashKey() != int2.HashKey() {
		t.Errorf("integers with same content have different hash keys")
	}

	if int1.HashKey() == int3.HashKey() {
		t.Errorf("integers with different content have same hash keys")
	}
}
