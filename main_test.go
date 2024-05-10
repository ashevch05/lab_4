package main

import "testing"

func TestValueTypeMod(t *testing.T) {
	x := 5
	value_type_modify(x)
	if x != 5 {
		t.Errorf("Expected x to be 5 after modifyValueType, got %d", x)
	}
}

func TestRefTypeMod(t *testing.T) {
	slice := []int{1, 2, 3}
	ref_type_modify(slice)
	expected := []int{4, 2, 3}
	if !equalSlices(slice, expected) {
		t.Errorf("Expected slice to be %v after modifyReferenceType, got %v", expected, slice)
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestGetValueType(t *testing.T) {
	y := get_value_type()
	if y != 10 {
		t.Errorf("Expected getValueTypeInstance to return 10, got %d", y)
	}
}

func TestGetRefType(t *testing.T) {
	anotherSlice := get_ref_type()
	expected := []int{5, 6, 7}
	if !equalSlices(anotherSlice, expected) {
		t.Errorf("Expected getReferenceTypeInstance to return %v, got %v", expected, anotherSlice)
	}
}
