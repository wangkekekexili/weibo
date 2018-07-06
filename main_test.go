package main

import "testing"

func TestRemoveTags(t *testing.T) {
	tests := []struct {
		s   string
		exp string
	}{
		{"Hello World", "Hello World"},
		{"Hello<br />World", "HelloWorld"},
		{"<p>Hello World</p>", "Hello World"},
	}
	for _, test := range tests {
		got := removeTags(test.s)
		if got != test.exp {
			t.Fatalf("result of removing tags from %q = %q; want %q", test.s, got, test.exp)
		}
	}
}
