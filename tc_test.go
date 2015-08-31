package tc

import (
	"strings"
	"testing"
)

func TestIsValid(t *testing.T) {
	ok, err := IsValid("17857715056", "ferhat", "elmas", 1988)
	if err != nil || !ok {
		t.Fatal(ok, err)
	}
}

func TestNotValidWithCorrectParams(t *testing.T) {
	ok, err := IsValid("17857715056", "ferhat", "elmas", 1989)
	if err != nil || ok {
		t.Fatal(ok, err)
	}
}

func TestNotValidWithWrongNumber(t *testing.T) {
	_, err := IsValid("1785771505", "ferhat", "elmas", 1988)
	if err == nil {
		t.Fatal("It should have failed but err is nil")
	} else if !strings.Contains(err.Error(), "11 characters") {
		t.Fatal("Doesn't contain require error message", err)
	}
}

func TestNotValidWithWrongYear(t *testing.T) {
	ok, err := IsValid("17857715056", "ferhat", "elmas", 0)
	if err != nil || ok {
		t.Fatal(ok, err)
	}
}
