package adt_test

import (
	"testing"

	adt "github.com/testingallthethings/008-stack"
)

func TestEmpty(t *testing.T) {
	s := adt.NewStack()

	if s.Empty() != true {
		t.Error("Stack was not empty")
	}
}

func TestNotEmpty(t *testing.T) {
	s := adt.NewStack()
	s.Add("Bob")

	if s.Empty() != false {
		t.Error("Stack was empty")
	}
}

func TestSizeZero(t *testing.T) {
	s := adt.NewStack()

	if s.Size() != 0 {
		t.Errorf("Expected zero elements found: %d", s.Size())
	}
}

func TestSizeOne(t *testing.T) {
	s := adt.NewStack()
	s.Add("Bob")

	if s.Size() != 1 {
		t.Error("Incorrect Size")
		t.Log("Expected: 1")
		t.Logf("Actual: %d", s.Size())
	}
}

func TestSizeThree(t *testing.T) {
	s := adt.NewStack()
	s.Add("Bob")
	s.Add("Dave")
	s.Add("Sarah")

	if s.Size() != 3 {
		t.Errorf("Expected 3, found %d", s.Size())
	}
}

func TestPopOne(t *testing.T) {
	s := adt.NewStack()
	s.Add("Bob")

	v := s.Pop()

	if v != "Bob" {
		t.Errorf("Expected Bob, found %s", v)
	}

	if s.Size() != 0 {
		t.Errorf("Expected size 0, found %d", s.Size())
	}
}

func TestPopTwo(t *testing.T) {
	s := adt.NewStack()
	s.Add("Bob")
	s.Add("Alice")

	v := s.Pop()

	if v != "Alice" {
		t.Errorf("Expected Alice, found %s", v)
	}

	v2 := s.Pop()

	if v2 != "Bob" {
		t.Errorf("Expected Bob, found %s", v2)
	}
}
