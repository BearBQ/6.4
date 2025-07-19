package main

import (
	"fmt"
	"io"
	"net/http"
)

func Sum(a int, b int) int {
	return a + b
}

func main() {

}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func FilterEven(nums []int) []int {
	var out []int
	for _, n := range nums {
		if n%2 == 0 {
			out = append(out, n)
		}
	}
	return out
}

func Discount(value int) (int, error) {
	switch {
	case value <= 0:
		return 0, fmt.Errorf("значение меньше 0")
	case value > 1000:
		return 10, nil
	default:
		return 0, nil
	}
}

func GetData(client *http.Client, url string) (string, error) {
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
}
