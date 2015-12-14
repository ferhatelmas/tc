package tc

import (
	"strings"
	"testing"
)

func TestIsValidFor(t *testing.T) {
	ok, err := IsValidFor("17857715056", "ferhat", "elmas", 1988)
	if err != nil || !ok {
		t.Fatal(ok, err)
	}
}

func TestNotValidWithCorrectParams(t *testing.T) {
	ok, err := IsValidFor("17857715056", "ferhat", "elmas", 1989)
	if err != nil || ok {
		t.Fatal(ok, err)
	}
}

func TestNotValidWithWrongNumber(t *testing.T) {
	_, err := IsValidFor("1785771505", "ferhat", "elmas", 1988)
	if err == nil {
		t.Fatal("It should have failed but err is nil")
	} else if !strings.Contains(err.Error(), "11 characters") {
		t.Fatal("Doesn't contain require error message", err)
	}
}

func TestNotValidWithWrongYear(t *testing.T) {
	ok, err := IsValidFor("17857715056", "ferhat", "elmas", 0)
	if err != nil || ok {
		t.Fatal(ok, err)
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		id    string
		valid bool
	}{
		{"10", false},
		{"111111111ab", false},
		{"10000000146", true}, // Ataturk
		{"10000000147", false},
		{"17857715056", true},
		{"17857715050", false},
	}
	for i, test := range tests {
		if IsValid(test.id) != test.valid {
			t.Errorf("%d: Expected %t, got %t for %s", i, test.valid, IsValid(test.id), test.id)
		}
	}
}
