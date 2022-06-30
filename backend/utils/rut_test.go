package utils

import (
	"testing"
)

func TestRutTypeRun(t *testing.T) {
	typeTest := Run
	want := "RUN"
	msg := typeTest.String()
	if msg != want {
		t.Fatalf(`typeTest.String() = %q, want match for %#q, nil`, msg, want)
	}
}

func TestRutTypeRut(t *testing.T) {
	typeTest := Rut
	want := "RUT"
	msg := typeTest.String()
	if msg != want {
		t.Fatalf(`typeTest.String() = %q, want match for %#q, nil`, msg, want)
	}
}

func TestRutTypeUnknown(t *testing.T) {
	var typeTest RutType = 2
	want := "unknown"
	msg := typeTest.String()
	if msg != want {
		t.Fatalf(`typeTest.String() = %q, want match for %#q, nil`, msg, want)
	}
}

func TestCleanRut(t *testing.T) {
	want := "8675309K"
	msg := "8.675.309-k"
	CleanRut(&msg)
	if msg != want {
		t.Fatalf(`CleanRut(&msg) = %q, want match for %#q, nil`, msg, want)
	}
}

func TestPrettyRut(t *testing.T) {
	want := "8.675.309-K"
	msg := "8675309k"
	PrettyRut(&msg)
	if msg != want {
		t.Fatalf(`PrettyRut(&msg) = %q, want match for %#q, nil`, msg, want)
	}
}

func TestGenerateVerifier(t *testing.T) {
	want := "K"
	msg, err := generateVerifier("8675309")
	if msg != want {
		t.Fatalf(`generateVerifier("8675309") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestGenerateVerifier2(t *testing.T) {
	want := "9"
	msg, err := generateVerifier("86753095")
	if msg != want {
		t.Fatalf(`generateVerifier("86753095") = %q, %v, want match for %q, nil`, msg, err, want)
	}
}

func TestGenerateVerifierInvalidString(t *testing.T) {
	want := "invalid RUT"
	_, err := generateVerifier("8675f309")
	if err.Error() != want {
		t.Fatalf(`generateVerifier("8675f309") = %q, want %q, nil`, err.Error(), want)
	}
}

func TestIsValidRut(t *testing.T) {
	want := true
	msg, err := IsValidRut("8675309K")
	if msg != want {
		t.Fatalf(`IsValidRut("8675309K") = false, %v, want match for true, nil`, err)
	}
}

func TestIsValidRutInvalid(t *testing.T) {
	want := false
	msg, err := IsValidRut("86753T99")
	if msg != want {
		t.Fatalf(`IsValidRut("86753T99") = true, %v, want match for false, nil`, err)
	}
}

func TestIsValidRutInvalidIdentifier(t *testing.T) {
	want := false
	msg, err := IsValidRut("8675309C")
	if msg != want {
		t.Fatalf(`IsValidRut("8675309C") = true, %v, want match for false, nil`, err)
	}
}

func TestIsValidRutIncorrectLength(t *testing.T) {
	want := false
	msg, err := IsValidRut("123234")
	if msg != want {
		t.Fatalf(`IsValidRut("123234") = true, %v, want match for false, nil`, err)
	}
}

func TestIsValidRutIncorrectIdentifier(t *testing.T) {
	want := false
	msg, err := IsValidRut("86753096")
	if msg != want {
		t.Fatalf(`IsValidRut("86753096") = true, %v, want match for false, nil`, err)
	}
}

func TestGetRutTypeRun(t *testing.T) {
	want := Run
	msg, err := GetRutType("8675309K")
	if msg != want {
		t.Fatalf(`GetRutType("8675309") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestGetRutTypeRut(t *testing.T) {
	want := Rut
	msg, err := GetRutType("867530959")
	if msg != want {
		t.Fatalf(`GetRutType("867530959") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestGetRutTypeRutInvalid(t *testing.T) {
	want := "invalid RUN/RUT"
	_, err := GetRutType("8675f309")
	if err.Error() != want {
		t.Fatalf(`GetRutType("8675f309") = %q, want %q, nil`, err.Error(), want)
	}
}
