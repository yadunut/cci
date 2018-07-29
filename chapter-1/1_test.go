package chapter1

import "testing"

func TestUniqueStr(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"", -1},
		{"abcdefg", 1},
		{"AAAAAAA", 0},
	}

	for _, c := range cases {
		got := UniqueStr(c.in)
		if got != c.want {
			t.Errorf("UniqueStr(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestBetterUniqueStr(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"", false},
		{"abcdefg", true},
		{"AAAAAAA", false},
	}

	for _, c := range cases {
		got := BetterUniqueStr(c.in)
		if got != c.want {
			t.Errorf("UniqueStr(%q) == %t, want %t", c.in, got, c.want)
		}
	}

}
