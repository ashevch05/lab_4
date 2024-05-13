package main

import (
	"runtime"
	"testing"
	"unsafe"
)

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

func TestValueTypeOnStack(t *testing.T) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	initialHeapSize := m.HeapInuse

	y := get_value_type()
	_ = y

	// Get the new heap size
	runtime.ReadMemStats(&m)
	newHeapSize := m.HeapInuse

	if newHeapSize-initialHeapSize != 0 {
		t.Errorf("Expected no heap allocation for value type, but heap grew by %d bytes", newHeapSize-initialHeapSize)
	}
}

func TestRefTypeOnHeap(t *testing.T) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	initialHeapSize := m.HeapInuse

	largeSlice := make([]int, 1000000)
	_ = largeSlice

	runtime.ReadMemStats(&m)
	newHeapSize := m.HeapInuse

	if newHeapSize <= initialHeapSize {
		t.Errorf("Expected heap allocation for large reference type, but heap size did not change")
	}
}

func TestGarbageCollection(t *testing.T) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	initialHeapSize := m.HeapInuse

	largeSlice := make([]int, 1000000)
	_ = largeSlice

	// Get the new heap size
	runtime.ReadMemStats(&m)
	heapSizeAfterAllocation := m.HeapInuse

	runtime.GC()
	runtime.GC()

	runtime.ReadMemStats(&m)
	heapSizeAfterGC := m.HeapInuse

	if heapSizeAfterGC >= heapSizeAfterAllocation || heapSizeAfterGC >= initialHeapSize {
		t.Errorf("Expected heap size to decrease after garbage collection, but it grew or remained the same")
	}
}

func TestSizeOfValueType(t *testing.T) {
	intSize := unsafe.Sizeof(int(0))
	if intSize != 8 {
		t.Errorf("Expected size of int to be 8 bytes, got %d bytes", intSize)
	}
}

func TestSizeOfRefType(t *testing.T) {
	sliceSize := unsafe.Sizeof([]int{})
	if sliceSize != 24 {
		t.Errorf("Expected size of slice to be 24 bytes, got %d bytes", sliceSize)
	}
}
