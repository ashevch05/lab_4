package main

import (
	"fmt"
	"unsafe"
)

func value_type_modify(x int) {
	x = 10
}

func ref_type_modify(slice []int) {
	slice[0] = 4
}

func get_value_type() int {
	y := 10
	return y
}

func get_ref_type() []int {
	anotherSlice := []int{5, 6, 7} // Reference type, stored on the heap
	return anotherSlice
}

// Garbage Collection example
type MyStruct struct {
	value int
}

func create_garbage_collection() *MyStruct {
	obj := &MyStruct{value: 42} // Reference type object on the heap
	fmt.Println("Object before collection:", obj)

	var obj2 *MyStruct = nil // Set the object to nil, eligible for garbage collection
	fmt.Println("Object set to nil, will be collected by garbage collector")

	return obj2 // Returning nil to demonstrate garbage collection
}

func main() {
	// Value type example
	x := 5
	value_type_modify(x)
	fmt.Println("\nValue type after modification:", x) // Output: 5

	// Reference type example
	slice := []int{1, 2, 3}
	ref_type_modify(slice)
	fmt.Println("Reference type after modification:", slice) // Output: [4 2 3]

	// Stack and Heap example
	y := get_value_type()
	fmt.Printf("Value type instance: %v (type: %T)\n", y, y) // Output: Value type instance: 10 (type: int)

	anotherSlice := get_ref_type()
	fmt.Printf("Reference type instance: %v (type: %T)\n", anotherSlice, anotherSlice) // Output: Reference type instance: [5 6 7] (type: []int)

	// Garbage Collection example
	_ = create_garbage_collection()

	// Print size of value type and reference type
	fmt.Println("Size of int:", unsafe.Sizeof(int(0)))    // Output: Size of int: 8
	fmt.Println("Size of slice:", unsafe.Sizeof([]int{})) // Output: Size of slice: 24
}
