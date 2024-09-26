package main

import "testing"

// func TestFactorial(t *testing.T) {
// 	result := Factorial(5)

// 	if result != 120 {
// 		t.Errorf("Result of Add is Wrong expect %d got %d", 120, result)
// 	}

// }

func TestFactorial(t *testing.T) {
	testCase := []struct {
		name     string
		num      int
		expected int
	}{
		{"Factorial 5", 5, 120},
		{"Factorial 2", 2, 2},
	}

	for _, tc := range testCase {
		// result := Factorial(tc.num)

		// if result != tc.expected {
		// 	t.Errorf("Result of Factorial is Wrong expect %d got %d", tc.expected, result)
		// }
		t.Run(tc.name, func(t *testing.T) {
			result := Factorial(tc.num)

			if result != tc.expected {
				t.Errorf("Result of Factorial is Wrong expect %d got %d", tc.expected, result)
			}
		})
	}

}

// func TestAdd(t *testing.T) {
// 	testCase := []struct {
// 		name     string
// 		a, b     int
// 		expected int
// 	}{
// 		{"Add + num", 2, 3, 5},
// 		{"Add - num", -1, -3, -4},
// 		{"Add 0", 0, 0, 0},
// 	}

// 	for _, tc := range testCase {
// 		// result := Add(tc.a, tc.b)

// 		// if result != tc.expected {
// 		// 	t.Errorf("Result of Add is Wrong expect %d got %d", tc.expected, result)
// 		// }
// 		t.Run(tc.name, func(t *testing.T) {
// 			result := Add(tc.a, tc.b)

// 			if result != tc.expected {
// 				t.Errorf("Result of Add is Wrong expect %d got %d", tc.expected, result)
// 			}
// 		})
// 	}

// }
