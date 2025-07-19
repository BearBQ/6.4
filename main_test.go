package main

import (
	"net/http"
	"net/http/httptest"
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

func TestDiscount(t *testing.T) {
	tests := []struct {
		value, discountValue int
	}{
		{0, 0},
		{1001, 10},
		{121, 0},
		{-25, 0},
	}
	for _, tt := range tests {
		output, err := Discount(tt.value)
		_ = err
		if output != tt.discountValue {
			t.Errorf("не совпадает")
		}
	}
}

func TestGetData(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/test-path" {
			t.Errorf("Ожидался запрос к /test-path, получен %s", r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if r.Method != http.MethodGet {
			t.Errorf("Ожидался GET-запрос, получен %s", r.Method)
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test response data"))
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := server.Client()
	url := server.URL + "/test-path"
	data, err := GetData(client, url)
	if err != nil {
		t.Errorf("ошибка")
	}
	expected := "test response data"
	if data != expected {
		t.Errorf("Ожидался ответ %q, получен %q", expected, data)
	}

}

func BenchmarkReverse(b *testing.B) {
	words := []string{"asdfffdsa", "aaaaaaaaaaaaa", "🙂🔥🙂🔥🙂🔥🙂🔥"}
	for _, w := range words {
		b.Run(w, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Reverse(w)
			}
		})
	}
}
