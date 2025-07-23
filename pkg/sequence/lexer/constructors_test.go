package lexer

import (
	"maps"
	"slices"
	"testing"
)

func TestLifeline(t *testing.T) {
	tcs := map[string]*LifelineDecl{
		"actor a as Alice":       {Type: "actor", Name: "a", Alias: "Alice"},
		"actor a":                {Type: "actor", Name: "a", Alias: ""},
		"participant a as Alice": {Type: "participant", Name: "a", Alias: "Alice"},
		"participant a":          {Type: "participant", Name: "a", Alias: ""},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(LifelineDecl{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (LifelineDecl{}).construct(input).(*LifelineDecl)
			if *got != *expected {
				t.Errorf("expected %#v got %#v", expected, got)
			}
		})
	}
}

func TestMessage(t *testing.T) {
	tcs := map[string]*Message{
		"a->>b":      {From: "a", To: "b", Content: ""},
		"a->>b: ACK": {From: "a", To: "b", Content: "ACK"},
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			if !(Message{}).check(input) {
				t.Errorf("checker has unexpectedly rejected the input")
			}
			expected := tcs[input]
			got := (Message{}).construct(input).(*Message)
			if *got != *expected {
				t.Errorf("expected %#v got %#v", expected, got)
			}
		})
	}
}
