package emojilib

import "testing"

func TestFind(t *testing.T) {
	for _, tt := range []struct {
		in   string
		want string
		err  error
	}{
		{"unknown", "", ErrUnknownEmoji},
		{"green_heart", "💚", nil},
		{"bee", "🐝", nil},
		{"scream", "😱", nil},
		{"rocket", "🚀", nil},
	} {
		got, err := Find(tt.in)
		if err != tt.err {
			t.Errorf("unexpected error")
		}

		if got.Char != tt.want {
			t.Errorf("Find(%q) = %q, nil, want %q, nil", tt.in, got.Char, tt.want)
		}
	}
}

func TestKeyword(t *testing.T) {
	for _, tt := range []struct {
		in   string
		want string
		err  error
	}{
		{"unknown", "", ErrUnknownKeyword},
		{"affection", "💙", nil},
		{"911", "🚑", nil},
		{"beef", "🐮", nil},
	} {
		got, err := Keyword(tt.in)
		if err != tt.err {
			t.Errorf("unexpected error")
		}

		if len(got) > 0 {
			if got[0].Char != tt.want {
				t.Errorf("Keyword(%q) = %q, nil, want %q, nil", tt.in, got[0].Char, tt.want)
			}
		}
	}
}

func TestAll(t *testing.T) {
	all := All()

	if len(all) == 0 {
		t.Fatalf("no emojis returned")
	}

	if len(all) != len(emojis) {
		t.Fatalf("unexpected number of emojis returned")
	}
}

func TestReplace(t *testing.T) {
	for _, tt := range []struct {
		in   string
		want string
	}{
		{"foo :green_heart:  bar", "foo 💚  bar"},
		{":sunny:  and :cloud: ", "☀️  and ☁️ "},
	} {
		if got := Replace(tt.in); got != tt.want {
			t.Errorf("Replace(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestReplaceWithPadding(t *testing.T) {
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
