package parser

import (
	"fmt"
	"testing"

	"github.com/ufukty/diagramer/pkg/sequence/ast"
)

func PrintDiagram(d *ast.Diagram) {
	fmt.Println("Lifelines:")
	for _, p := range d.Lifelines {
		fmt.Printf("  %s (%s)\n", p.Name, p.Alias)
	}
	fmt.Println("Messages:")
	for _, m := range d.Messages {
		fmt.Printf("  %s ->> %s: %s\n", m.From, m.To, m.Content)
	}
}

func TestFile(t *testing.T) {
	diagram, err := File("testdata/1.txt")
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
	db, ok := diagram.Lifelines["db"]
	if !ok {
		t.Errorf("Expected lifelines 'db' to exist")
	} else if db.Alias != "Server Database" {
		t.Errorf("Expected db name 'Server Database', got '%s'", db.Alias)
	}

	u, ok := diagram.Lifelines["u"]
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
