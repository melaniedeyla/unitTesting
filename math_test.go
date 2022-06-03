package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 3)

	if result != 4 {
		t.Errorf("Add(1, 3) FAILED. Expected %d, got %d\n", 4, result)
	} else {
		t.Logf("Add(1, 3) PASSED. EXPECTED %d, got %d\n", 4, result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)

	if result != 2 {
		t.Errorf("Subtract(5, 3) FAILED. Expected %d, got %d\n", 2, result)
	} else {
		t.Logf("Subtract(5, 3) PASSED. Expected %d, got %d\n", 2, result)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(9, 3)

	if result != 3 {
		t.Errorf("Divide(9, 3) FAILED. Expected %d, got %f", 3, result)
	} else {
		t.Logf("Divide(9, 3) PASSED. Expected %d, got %f", 3, result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(2, 4)

	if result != 8 {
		t.Errorf("Multiply(2, 4) FAILED. Expected %d, got %d", 8, result)
	} else {
		t.Logf("Multiply(2, 4) PASSED. Expected %d, got %d", 8, result)
	}
}
