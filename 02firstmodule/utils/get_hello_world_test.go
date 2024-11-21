package utils

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, World from utils!"
	if got := GetHello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
