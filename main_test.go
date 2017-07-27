package main

import "testing"

func TestBacktick(t *testing.T) {
	var cases = []struct {
		s, e string
	}{
		{"a", "`a`"},
		{"ag", "`ag`"},
		{"```", "\"```\""},
		{"a`a-b=c`", "`a`+\n\"`\"+\n`a-b=c`+\n\"`\""},
	}
	for _, c := range cases {
		g := backtick(c.s)
		if g != c.e {
			t.Errorf("\nOrigin:\n%s\nExpected:\n%s\nGot:\n%s\n", c.s, c.e, g)
		}
	}
}
