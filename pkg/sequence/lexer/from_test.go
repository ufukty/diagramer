package lexer

import (
	"fmt"
	"os"
	"slices"
	"testing"
)

func PrintDiagram(d *Diagram) {
	fmt.Println("Lifelines:")
	for _, p := range d.Lifelines {
		fmt.Printf("  %s (%s)\n", p.Name, p.Alias)
	}
	fmt.Println("Messages:")
	for _, m := range d.Messages {
		fmt.Printf("  %s ->> %s: %s\n", m.From, m.To, m.Content)
	}
}

func findLifeline(diagram *Diagram, lifeline string) (*Lifeline, bool) {
	i := slices.IndexFunc(diagram.Lifelines, func(ll *Lifeline) bool { return ll.Name == lifeline })
	if i != -1 {
		return diagram.Lifelines[i], true
	}
	return nil, false
}

func TestFromReader(t *testing.T) {
	file, err := os.Open("testdata/1.txt")
	if err != nil {
		t.Fatalf("prep, open file: %v", err)
	}
	defer file.Close() //nolint:errcheck
	diagram, err := FromReader(file)
	if err != nil {
		t.Fatalf("act: %v", err)
	}

	// Check autoNumber
	if !diagram.AutoNumber {
		t.Errorf("Expected AutoNumber to be true")
	}

	// Check lifelines count
	if len(diagram.Lifelines) != 7 {
		t.Errorf("Expected 7 lifeliness, got %d", len(diagram.Lifelines))
	}

	// Check specific lifelines names

	db, ok := findLifeline(diagram, "db")
	if !ok {
		t.Errorf("Expected lifelines 'db' to exist")
	} else if db.Alias != "Server Database" {
		t.Errorf("Expected db name 'Server Database', got '%s'", db.Alias)
	}

	u, ok := findLifeline(diagram, "u")
	if !ok {
		t.Errorf("Expected lifelines 'u' to exist")
	} else if u.Alias != "User" {
		t.Errorf("Expected 'u' name 'User', got '%s'", u.Alias)
	}

	// Check number of messages
	expectedMsgs := 17
	if len(diagram.Messages) != expectedMsgs {
		t.Errorf("Expected %d messages, got %d", expectedMsgs, len(diagram.Messages))
	}

	// Check first message
	firstMsg := diagram.Messages[0]
	if firstMsg.From != "u" || firstMsg.To != "a" || firstMsg.Content != "opens the website" {
		t.Errorf("First message parsed incorrectly: %+v", firstMsg)
	}

	PrintDiagram(diagram)
}
