package emojilib

import "testing"

func TestReplacer(t *testing.T) {
	for _, tt := range []struct {
		in   string
		want string
	}{
		{"foo :green_heart: bar", "foo 💚  bar"},
		{":sunny:", "☀️ "},
	} {
		if got := ReplaceWithPadding(tt.in); got != tt.want {
			t.Errorf("ReplaceWithPadding(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
