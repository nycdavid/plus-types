package main

type Enumerable struct {
	elements []string
}

func (enum *Enumerable) each(block func(string)) {
	for _, element := range enum.elements {
		block(element)
	}
}

func (enum *Enumerable) append(element string) {
	enum.elements = append(enum.elements, element)
}