package requestcorrelation

import (
	"strings"
	"testing"
)

func TestGenerateRandomHexString(t *testing.T) {
	s := GenerateRandomHexString(10)
	if len(s) != 10 {
		t.Error("s length is not 10")
	}

	const hexAlphabet = "0123456789abcdef"

	for _, r := range s {
		if !strings.ContainsRune(hexAlphabet, r) {
			t.Errorf("Rune %s was found in string and is invalid", string(r))
		}
	}
}

func TestGenerateNewTraceParent(t *testing.T) {
	tp := GenerateNewTraceParent()

	if tp.Version != "00" {
		t.Error("Version should be 00")
	}

	if tp.Flags != "00" {
		t.Error("Flags should be 00")
	}

	if len(tp.TraceID) != 32 {
		t.Error("TraceID length should be 32")
	}

	if len(tp.SpanID) != 16 {
		t.Error("SpanID length should be 16")
	}
}

func TestNewSpanFromParent(t *testing.T) {
	parent := GenerateNewTraceParent()
	child := NewSpanFromParent(parent)

	if child.Version != parent.Version {
		t.Error("Child Version should match parent Version")
	}

	if child.Flags != parent.Flags {
		t.Error("Child Flags should match parent Flags")
	}

	if child.TraceID != parent.TraceID {
		t.Error("Child TraceID should match parent TraceID")
	}

	if child.SpanID == parent.SpanID {
		t.Error("Child SpanID should not match parent SpanID")
	}

}

func TestNewTraceParentValid(t *testing.T) {
	tp, err := NewTraceParent(0, "12345678901234567890123456789012", "1234567890123456", 0)
	if err != nil {
		t.Errorf("Error parsing trace parent. Error: %v", err)
	}

	if tp.Version != "00" {
		t.Errorf("Child Version should be 00, but was %v", tp.Version)
	}

	if tp.TraceID != "12345678901234567890123456789012" {
		t.Errorf("Child TraceID should be 12345678901234567890123456789012, but was %v", tp.TraceID)
	}

	if tp.SpanID != "1234567890123456" {
		t.Errorf("Child SpanID should be 1234567890123456, but was %v", tp.SpanID)
	}

	if tp.Flags != "00" {
		t.Errorf("Child Flags should be 00, but was %v", tp.Flags)
	}

}

func TestNewTraceParentInvalid(t *testing.T) {
	_, err := NewTraceParent(0, "", "", 0)
	if err == nil {
		t.Errorf("Expected an error parsing trace parent. Error: %v", err)
	}
}

func TestNewTraceParentWithStrValid(t *testing.T) {
	tp, err := NewTraceParentWithStr("00", "12345678901234567890123456789012", "1234567890123456", "00")
	if err != nil {
		t.Errorf("Error parsing trace parent. Error: %v", err)
	}

	if tp.Version != "00" {
		t.Errorf("Child Version should be 00, but was %v", tp.Version)
	}

	if tp.TraceID != "12345678901234567890123456789012" {
		t.Errorf("Child TraceID should be 12345678901234567890123456789012, but was %v", tp.TraceID)
	}

	if tp.SpanID != "1234567890123456" {
		t.Errorf("Child SpanID should be 1234567890123456, but was %v", tp.SpanID)
	}

	if tp.Flags != "00" {
		t.Errorf("Child Flags should be 00, but was %v", tp.Flags)
	}

}
