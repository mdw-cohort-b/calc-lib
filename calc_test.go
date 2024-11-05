package calc

import (
	"fmt"
	"testing"
)

func TestAddition_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 0, b: 0, want: 0},
		{a: 0, b: 1, want: 1},
		{a: 1, b: 1, want: 2},
		{a: 2, b: 3, want: 5},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d + %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			addition := &Addition{}
			got := addition.Calculate(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestSubtraction_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 0, b: 0, want: 0},
		{a: 0, b: 1, want: -1},
		{a: 1, b: 1, want: 0},
		{a: 5, b: 3, want: 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d - %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			subtraction := &Subtraction{}
			got := subtraction.Calculate(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMultiplication_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 0, b: 0, want: 0},
		{a: 1, b: 1, want: 1},
		{a: 4, b: 3, want: 12},
		{a: 5, b: 3, want: 15},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d * %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			multiplication := &Multiplication{}
			got := multiplication.Calculate(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestDivision_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 1, b: 1, want: 1},
		{a: 12, b: 3, want: 4},
		{a: 11, b: 2, want: 5},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d * %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			division := &Division{}
			got := division.Calculate(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestDivision_ByZeroPanics(t *testing.T) {
	division := &Division{}
	defer func() {
		r := recover()
		t.Log(r)
		if r == nil {
			t.Fatal("expected panic")
		}
	}()
	division.Calculate(1, 0)
}
