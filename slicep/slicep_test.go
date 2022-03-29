package slicep

import (
	"testing"
)

func TestMap_WithString(t *testing.T) {
	a := &Slicep{Arry: []interface{}{"foo", "bar"}}
	b := a.Map(func(i int, el interface{}) interface{} {
		return el.(string) + "a"
	})

	if b.At(0) != "fooa" {
		t.Errorf("Expected fooa, got %s", b.At(0))
	}
	if b.At(1) != "bara" {
		t.Errorf("Expected bara, got %s", b.At(1))
	}
}

func TestMap_WithInt(t *testing.T) {
	a := &Slicep{Arry: []interface{}{0, 1, 2, 3}}
	b := a.Map(func(i int, el interface{}) interface{} {
		return el.(int) + 1
	})

	if b.At(0) != 1 {
		t.Errorf("Expected 1, got %d", b.At(0))
	}
	if b.At(1) != 2 {
		t.Errorf("Expected 2, got %d", b.At(1))
	}
}

func TestSum_WithInt(t *testing.T) {
	a := &Slicep{Arry: []interface{}{0, 1, 2, 3}}

	sum := a.Sum()

	if sum != 6 {
		t.Errorf("Expected 6, got %d", sum)
	}
}

func TestLength(t *testing.T) {
	a := &Slicep{Arry: []interface{}{0, 1, 2, 3}}

	if a.Length() != 4 {
		t.Errorf("Expected 4, got %d", a.Length())
	}
}

func TestSelect_WithString(t *testing.T) {
	a := &Slicep{Arry: []interface{}{"foo", "bar", "baz"}}
	b := a.Select(func(i int, el interface{}) interface{} {
		return el.(string) == "bar"
	})

	if b.Length() != 1 {
		t.Errorf("Expected 1, got %d", b.Length())
	}
	if b.At(0) != "bar" {
		t.Errorf("Expected bar, got %s", b.At(0))
	}
}

func TestIndex_WithInt(t *testing.T) {
	s := &Slicep{Arry: []interface{}{1, 2, 3, 4}}

	i, _ := s.Index(3)

	if i != 2 {
		t.Errorf("Expected %d but got %d", 2, i)
	}
}

func TestIndexErr_WithInt(t *testing.T) {
	s := &Slicep{Arry: []interface{}{1, 2, 3, 4}}

	_, err := s.Index(5)

	if err == nil {
		t.Errorf("Expected error for non-existent element")
	}
}

func TestReverse_WithInt(t *testing.T) {
	s := &Slicep{Arry: []interface{}{1, 2, 3, 4}}

	ns := s.Reverse()

	if ns.At(0) != 4 {
		t.Errorf("Expected first element to be 4, got %d", ns.At(0))
	}
}

func TestReverse_WithString(t *testing.T) {
	s := &Slicep{Arry: []interface{}{"foo", "bar", "baz"}}

	ns := s.Reverse()

	if ns.At(0) != "baz" {
		t.Errorf("Expected first element to be baz, got %s", ns.At(0))
	}
}

func TestConcat_WithString(t *testing.T) {
	s1 := &Slicep{Arry: []interface{}{"foo"}}
	s2 := &Slicep{Arry: []interface{}{"bar"}}

	ns := s1.Concat(s2)

	expectedLength := 2

	if ns.Length != expectedLength {

	}
}
