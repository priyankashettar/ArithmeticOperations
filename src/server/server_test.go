package main

import "testing"

// Unit Testcases for the basic Math Operations
func TestAdd(t *testing.T) {
	sum := Add(2, 3)
	expected := 5
	if sum != float64(expected) {
		t.Errorf("Sum is : %v but expected value is: %v", sum, expected)
	}
}

func TestSubtract(t *testing.T) {
	subtractResult := Subtract(4, 2)
	expected := 2
	if subtractResult != float64(expected) {
		t.Errorf("Subtraction result is : %v but expected value is : %v", subtractResult, expected)
	}
}

func TestMultiply(t *testing.T) {
	multiplicationResult := Multiply(3, 3)
	expected := 9
	if multiplicationResult != float64(expected) {
		t.Errorf("Multiplication result is : %v but expected value is %v", multiplicationResult, expected)
	}
}

func TestDivide(t *testing.T) {
	divisionResult := Division(8, 2)
	expected := 4
	if divisionResult != float64(expected) {
		t.Errorf("Division result is : %v but expected value is %v", divisionResult, expected)
	}
}
