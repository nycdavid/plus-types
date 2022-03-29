package slicep

import "errors"

type Slicep struct {
	Arry []interface{}
}

func (sp *Slicep) Map(cb func(i int, el interface{}) interface{}) *Slicep {
	newSlicep := &Slicep{}

	for i, el := range sp.Arry {
		new_el := cb(i, el)
		newSlicep.Arry = append(newSlicep.Arry, new_el)
	}

	return newSlicep
}

func (sp *Slicep) At(idx int) interface{} {
	return sp.Arry[idx]
}

func (sp *Slicep) Sum() int {
	sum := 0

	for _, el := range sp.Arry {
		sum += el.(int)
	}

	return sum
}

func (sp *Slicep) Length() int {
	return len(sp.Arry)
}

func (sp *Slicep) Select(cb func(i int, el interface{}) interface{}) *Slicep {
	newSp := &Slicep{}

	for i, el := range sp.Arry {
		if cb(i, el).(bool) {
			newSp.Arry = append(newSp.Arry, el)
		}
	}

	return newSp
}

func (sp *Slicep) Index(el interface{}) (int, error) {
	for i, _el := range sp.Arry {
		if _el == el {
			return i, nil
		}
	}

	return -1, errors.New("Element not found")
}

func (sp *Slicep) Reverse() *Slicep {
	a := make([]interface{}, len(sp.Arry))
	i := len(sp.Arry) - 1

	for _, el := range sp.Arry {
		a[i] = el
		i--
	}

	return &Slicep{Arry: a}
}
