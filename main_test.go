package main

import "testing"

func TestStripHTML(t *testing.T) {
	tests := []struct {
		in, exp string
	}{
		{"Hey There<br />", "Hey There"},
		{`Hello<a href=""><span>Hey There</span></a>World`, "HelloWorld"},
		{"今天<span >中华儿女</span>", "今天"},
	}
	for _, test := range tests {
		got := stripHTML(test.in)
		if got != test.exp {
			t.Fatalf("got %q; want %q", got, test.exp)
		}
	}
}
