package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloNames(t *testing.T) {
	isGreeting := func(name string) func(greeting string) bool {
		want := regexp.MustCompile(`\b` + name + `\b`)
		return want.MatchString
	}

	names := []string{
		"Someone",
		"SomeoneElse",
	}

	msgs, err := Hellos(names)
	if err != nil {
		t.Error("Hellos(names) failed!")
	}

	for name, msg := range msgs {
		if !isGreeting(name)(msg) {
			t.Errorf(`Hello("%q") = %q, %v`, name, msg, err)
		}
	}
}

func TestHelloNamesEmpty(t *testing.T) {
	names := []string{
		"",
		"SomeoneElse",
	}

	_, err := Hellos(names)
	if err == nil {
		t.Error("Hellos(names) succesful even though name is empty!")
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
