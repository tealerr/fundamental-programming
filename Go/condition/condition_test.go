package condition

import (
	"testing"
)

func TestWhenInputScore_80_ShouldReturn_A(t *testing.T) {
	// Test case 1: Score is 80, should return "A"
	result := CalculateGrade(80)
	expected := 'A'
	if result != expected {
		t.Errorf("Expected %c but got %c", expected, result)
	}
}

func TestWhenInputScore_74_ShouldReturn_B(t *testing.T) {
	// Test case 2: Score is 74, should return "B"
	result := CalculateGrade(74)
	expected := 'B'
	if result != expected {
		t.Errorf("Expected %c but got %c", expected, result)
	}
}

func TestWhenInputOtherScore_ShouldReturn_F(t *testing.T) {
	// Test case 3: Score is not 80, should return "F"
	result := CalculateGrade(62)
	expected := 'F'
	if result != expected {
		t.Errorf("Expected %c but got %c", expected, result)
	}
}

func TestWhenGradeIs_A_shouldReturnMessageGrade_A(t *testing.T) {
	result := ShowGrade('A')
	expected := "Your grade is A"

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestWhenGradeIs_B_shouldReturnMessageGrade_B(t *testing.T) {
	result := ShowGrade('B')
	expected := "Your grade is B"

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestWhenGradeIs_F_shouldReturnMessageGrade_F(t *testing.T) {
	result := ShowGrade('F')
	expected := "Your grade is F"

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestWhenInput_0_should_Return_DeviceIsOff(t *testing.T) {
	result := IsActive(0)
	expected := "Device is Off"

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestWhenInput_1_should_Return_DeviceIsOn(t *testing.T) {
	result := IsActive(1)
	expected := "Device is On"

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
