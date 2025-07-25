package main

import "testing"

func TestAddWithDataset(t *testing.T) {
	// Definisikan struktur untuk test case
	type testCase struct {
		a        int
		b        int
		expected int
		name     string // Nama deskriptif untuk setiap test case
	}

	// Buat slice dari testCase, ini adalah dataset kita
	tests := []testCase{
		{a: 2, b: 3, expected: 5, name: "Positive numbers"},
		{a: -2, b: -3, expected: -5, name: "Negative numbers"},
		{a: 5, b: -3, expected: 2, name: "Positive and negative numbers"},
		{a: 0, b: 0, expected: 0, name: "Zero input"},
		{a: 100, b: 200, expected: 300, name: "Larger positive numbers"},
		{a: -10, b: 5, expected: -5, name: "Negative and positive with negative result"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sum := Add(tc.a, tc.b)
			if sum != tc.expected {
				t.Errorf("Add(%d, %d) failed, got %d, expected %d", tc.a, tc.b, sum, tc.expected)
			}
		})
	}
}
