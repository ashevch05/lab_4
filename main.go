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
	anotherSlice := []int{5, 6, 7}
	return anotherSlice
}

type MyStruct struct {
	value int
}

func create_garbage_collection() *MyStruct {
	obj := &MyStruct{value: 42}
	fmt.Println("Object before collection:", obj)

	var obj2 *MyStruct = nil
	fmt.Println("Object set to nil, will be collected by garbage collector")

	return obj2
}

func main() {
	x := 5
	value_type_modify(x)
	fmt.Println("\nValue type after modification:", x)

	slice := []int{1, 2, 3}
	ref_type_modify(slice)
	fmt.Println("Reference type after modification:", slice)

	y := get_value_type()
	fmt.Printf("Value type instance: %v (type: %T)\n", y, y)

	anotherSlice := get_ref_type()
	fmt.Printf("Reference type instance: %v (type: %T)\n", anotherSlice, anotherSlice)

	_ = create_garbage_collection()

	fmt.Println("Size of int:", unsafe.Sizeof(int(0)))
	fmt.Println("Size of slice:", unsafe.Sizeof([]int{}))
}
