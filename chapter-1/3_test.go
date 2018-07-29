package chapter1

import "testing"

type input struct {
	in   struct{ a, b string }
	want bool
}

var cases = []input{
	{in: struct{ a, b string }{"", ""}, want: true},
	{in: struct{ a, b string }{"hello", "hellos"}, want: false},
	{in: struct{ a, b string }{"hello", "hello"}, want: true},
	{in: struct{ a, b string }{"ellho", "hello"}, want: true},
	{in: struct{ a, b string }{"ellho", "helll"}, want: false},
}

func TestPermutateStr(t *testing.T) {
	for _, c := range cases {
		got := PermutateStr(c.in.a, c.in.b)
		if got != c.want {
			t.Errorf("PermutateStr(%q, %q) == %t, want %t", c.in.a, c.in.b, got, c.want)
		}
	}
}
