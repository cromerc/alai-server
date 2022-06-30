package utils

import (
	"testing"
)

func TestReverse(t *testing.T) {
	want := "8675309"
	msg := Reverse("9035768")
	if msg != want {
		t.Fatalf(`Reverse("9035768") = %q, want match for %#q, nil`, msg, want)
	}
}

func TestInsertNth(t *testing.T) {
	want := "867.530.9"
	msg := InsertNth("8675309", 3, '.')
	if msg != want {
		t.Fatalf(`InsertNth("8675309") = %q, want match for %#q, nil`, msg, want)
	}
}

func TestToInt(t *testing.T) {
	want := 123
	msg := toInt("123")
	if msg != want {
		t.Fatalf(`toInt("123") = %q, want match for %#q, nil`, msg, want)
	}
}

func TestToIntInvalid(t *testing.T) {
	want := 123
	msg := toInt("12f3")
	if msg == want {
		t.Fatalf(`toInt("12f3") = %q, want match for %#q, nil`, msg, want)
	}
}
