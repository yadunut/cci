package chapter1

import "testing"

func TestEscapeSpace(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{""}, ""},
		{"", args{"Mr John Smith"}, "Mr%20John%20Smith"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EscapeSpace(tt.args.a); got != tt.want {
				t.Errorf("EscapeSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
