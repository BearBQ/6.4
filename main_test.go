package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 8, 9},
		{-1111111111, 1111111111, 0},
		{0, 0, 0},
		{-1, -5, -6},
	}
	for _, tt := range tests {
		if got := Sum(tt.a, tt.b); tt.expected != got {
			t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		word, result string
	}{
		{"жопа", "апож"},
		{"заказ", "заказ"},
		{"🙂🔥", "🔥🙂"},
	}
	for _, tt := range tests {
		if Reverse(tt.word) != tt.result {
			t.Errorf("Rewers для \"%s\" дал неверный результат", tt.word)
		}
	}
}

func TestFilterEven(t *testing.T) {
	tests := []struct {
		incoming, out []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 4}},
		{[]int{}, nil},
		{[]int{2, 4, 6}, []int{2, 4, 6}},
		{[]int{1, 3, 5}, nil},
	}
	for _, tt := range tests {
		result := FilterEven(tt.incoming)
		if !reflect.DeepEqual(result, tt.out) {
			t.Errorf("неверно работает. Вход: %v, выход: %v, ожидание: %v\n", tt.incoming, result, tt.out)
		}
	}
}
