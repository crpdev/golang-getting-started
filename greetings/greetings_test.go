package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Rajapandian"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Rajapandian")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Rajapandian) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
