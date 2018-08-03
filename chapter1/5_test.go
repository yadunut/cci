package chapter1

import "testing"

func TestCompressString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{"aaaa"}, want: "a4"},
		{args: args{"bbbb"}, want: "b4"},
		{args: args{"cccc"}, want: "c4"},
		{args: args{"dddd"}, want: "d4"},
		{args: args{"aaaabbbbccccdddd"}, want: "a4b4c4d4"},
		{args: args{"aabbccdd"}, want: "aabbccdd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompressString(tt.args.s); got != tt.want {
				t.Errorf("CompressString(\"%v\") = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}