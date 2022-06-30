package utils

import (
	"reflect"
	"testing"
	"time"
)

func TestDateUnmarshalJSON(t *testing.T) {
	want := "1985-02-23 00:00:00 +0000 UTC"
	var date Date
	err := date.UnmarshalJSON([]byte("\"1985-02-23\""))
	msg := time.Time(date).String()
	if msg != want {
		t.Fatalf(`date.UnmarshalJSON([]byte("\"1985-02-23\"") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestDateUnmarshalJSONMalformed(t *testing.T) {
	want := error(&time.ParseError{})
	var date Date
	err := date.UnmarshalJSON([]byte("\"1985/02/23\""))
	msg := time.Time(date).String()
	if reflect.TypeOf(err) != reflect.TypeOf(want) {
		t.Fatalf(`date.UnmarshalJSON([]byte("\"1985/02/23\"") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestDateMarshalJSON(t *testing.T) {
	want := "\"1985-02-23T00:00:00Z\""
	parsed, _ := time.Parse("2006-01-02", "1985-02-23")
	var date Date = Date(parsed)
	msg, err := date.MarshalJSON()
	if string(msg) != want {
		t.Fatalf(`date.MarshalJSON() = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestDateFormat(t *testing.T) {
	want := "Feb 23, 1985"
	parsed, _ := time.Parse("2006-01-02", "1985-02-23")
	var date Date = Date(parsed)
	msg := date.Format("Jan 2, 2006")
	if msg != want {
		t.Fatalf(`date.Format("Jan 2, 2006") = %q, want match for %#q, nil`, msg, want)
	}
}

func TestDateScan(t *testing.T) {
	want := "0001-01-01T00:00:00Z"
	parsed, _ := time.Parse("2006-01-02", "1985-02-23")
	var date Date = Date(parsed)
	err := date.Scan(Date{})
	msg := date.Format("2006-01-02T15:04:05Z07:00")
	if msg != want {
		t.Fatalf(`date.Scan(Date{}) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestDateValue(t *testing.T) {
	var want time.Time
	parsed, _ := time.Parse("2006-01-02", "1985-02-23")
	var date Date = Date(parsed)
	msg, _ := date.Value()
	if reflect.TypeOf(msg) != reflect.TypeOf(want) {
		t.Fatalf(`date.Value() = %q, want match for %#q, nil`, msg, want)
	}
}

func TestDateGormDateType(t *testing.T) {
	want := "date"
	parsed, _ := time.Parse("2006-01-02", "1985-02-23")
	var date Date = Date(parsed)
	msg := date.GormDataType()
	if msg != want {
		t.Fatalf(`date.GormDateType() = %q, want match for %#q, nil`, msg, want)
	}
}

func TestDateGobEncode(t *testing.T) {
	want := "\x01\x00\x00\x00\x0e\x94\x0f!\x00\x00\x00\x00\x00\xff\xff"
	parsed, _ := time.Parse("2006-01-02", "1985-02-23")
	var date Date = Date(parsed)
	msg, err := date.GobEncode()
	if string(msg) != want {
		t.Fatalf(`date.GobEncode() = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestDateGobDecode(t *testing.T) {
	want := "1985-02-23T00:00:00Z"
	parsed, _ := time.Parse("2006-01-02", "2006-01-02")
	var date Date = Date(parsed)
	err := date.GobDecode([]byte("\x01\x00\x00\x00\x0e\x94\x0f!\x00\x00\x00\x00\x00\xff\xff"))
	msg := date.Format("2006-01-02T15:04:05Z07:00")
	if string(msg) != want {
		t.Fatalf(`date.GobDecode([]byte()) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestDateTimeUnmarshalJSON(t *testing.T) {
	want := "1985-02-23 12:13:14 +0000 UTC"
	var dateTime DateTime
	err := dateTime.UnmarshalJSON([]byte("\"1985-02-23 12:13:14\""))
	msg := time.Time(dateTime).String()
	if msg != want {
		t.Fatalf(`dateTime.UnmarshalJSON([]byte("\"1985-02-23 12:13:14\"") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestDateTimeUnmarshalJSONMalformed(t *testing.T) {
	want := error(&time.ParseError{})
	var dateTime DateTime
	err := dateTime.UnmarshalJSON([]byte("\"1985/02/23 12-13-14\""))
	msg := time.Time(dateTime).String()
	if reflect.TypeOf(err) != reflect.TypeOf(want) {
		t.Fatalf(`dateTime.UnmarshalJSON([]byte("\"1985/02/23 12-13-14\"") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestDateTimeMarshalJSON(t *testing.T) {
	want := "\"1985-02-23T12:13:14Z\""
	parsed, _ := time.Parse("2006-01-02 15:04:05", "1985-02-23 12:13:14")
	var dateTime DateTime = DateTime(parsed)
	msg, err := dateTime.MarshalJSON()
	if string(msg) != want {
		t.Fatalf(`dateTime.MarshalJSON() = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestDateTimeFormat(t *testing.T) {
	want := "Feb 23, 1985 12:13:14"
	parsed, _ := time.Parse("2006-01-02 15:04:05", "1985-02-23 12:13:14")
	var dateTime DateTime = DateTime(parsed)
	msg := dateTime.Format("Jan 2, 2006 15:04:05")
	if msg != want {
		t.Fatalf(`dateTime.Format("Jan 2, 2006 15:04:05") = %q, want match for %#q, nil`, msg, want)
	}
}

func TestDateTimeScan(t *testing.T) {
	want := "0001-01-01T00:00:00Z"
	parsed, _ := time.Parse("2006-01-02 15:04:05", "1985-02-23 12:13:14")
	var dateTime DateTime = DateTime(parsed)
	err := dateTime.Scan(Date{})
	msg := dateTime.Format("2006-01-02T15:04:05Z07:00")
	if msg != want {
		t.Fatalf(`dateTime.Scan(Date{}) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestDateTimeValue(t *testing.T) {
	var want time.Time
	parsed, _ := time.Parse("2006-01-02 15:04:05", "1985-02-23 12:13:14")
	var dateTime DateTime = DateTime(parsed)
	msg, _ := dateTime.Value()
	if reflect.TypeOf(msg) != reflect.TypeOf(want) {
		t.Fatalf(`dateTime.Value() = %q, want match for %#q, nil`, msg, want)
	}
}

func TestDateTimeGormDateType(t *testing.T) {
	want := "datetime"
	parsed, _ := time.Parse("2006-01-02 15:04:05", "1985-02-23 12:13:14")
	var dateTime DateTime = DateTime(parsed)
	msg := dateTime.GormDataType()
	if msg != want {
		t.Fatalf(`dateTime.GormDateType() = %q, want match for %#q, nil`, msg, want)
	}
}

func TestDateTimeGobEncode(t *testing.T) {
	want := "\x01\x00\x00\x00\x0e\x94\x0f\xcc\xda\x00\x00\x00\x00\xff\xff"
	parsed, _ := time.Parse("2006-01-02 15:04:05", "1985-02-23 12:13:14")
	var dateTime DateTime = DateTime(parsed)
	msg, err := dateTime.GobEncode()
	if string(msg) != want {
		t.Fatalf(`dateTime.GobEncode() = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestDateTimeGobDecode(t *testing.T) {
	want := "1985-02-23T12:13:14Z"
	parsed, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 12:13:14")
	var dateTime DateTime = DateTime(parsed)
	err := dateTime.GobDecode([]byte("\x01\x00\x00\x00\x0e\x94\x0f\xcc\xda\x00\x00\x00\x00\xff\xff"))
	msg := dateTime.Format("2006-01-02T15:04:05Z07:00")
	if string(msg) != want {
		t.Fatalf(`dateTime.GobDecode([]byte()) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
