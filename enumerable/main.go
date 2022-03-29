package main

import "fmt"

func main() {
	arry := &Enumerable{}
	arry.append("a")
	arry.append("b")
	arry.append("c")

	arry.each(func(element interface{}) {
		fmt.Println(element.(string))
	})

	has_a := arry.any(func(element interface{}) bool {
		return element == "a"
	})

	has_z := arry.any(func(element interface{}) bool {
		return element == "z"
	})

	fmt.Println(has_a)
	fmt.Println(has_z)
}

type Enumerable struct {
	elements []interface{}
}

func (enum *Enumerable) each(block func(interface{})) {
	for _, element := range enum.elements {
		block(element)
	}
}

func (enum *Enumerable) append(element interface{}) {
	enum.elements = append(enum.elements, element)
}

func (enum *Enumerable) any(block func(interface{}) bool) bool {
	v := false
	for _, element := range enum.elements {
		v = block(element)

		if v {
			break
		}
	}

	return v
}