package lexer

import (
	"fmt"
	"os"
	"slices"
	"testing"
)

func PrintDiagram(d *Diagram) {
	for _, stmt := range d.Stmts {
		fmt.Printf("%#v\n", stmt)
	}
}

func findLifeline(diagram *Diagram, lifeline string) (*LifelineDecl, bool) {
	i := slices.IndexFunc(diagram.Stmts, func(s Stmt) bool {
		ll, ok := s.(*LifelineDecl)
		if !ok {
			return false
		}
		return ll.Name == lifeline
	})
	if i != -1 {
		return diagram.Stmts[i].(*LifelineDecl), true
	}
	return nil, false
}

func findMessage(diagram *Diagram, from, to, content string) (*Message, bool) {
	i := slices.IndexFunc(diagram.Stmts, func(s Stmt) bool {
		m, ok := s.(*Message)
		if !ok {
			return false
		}
		return *m == Message{From: from, To: to, Content: content}
	})
	if i != -1 {
		return diagram.Stmts[i].(*Message), true
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
	if !diagram.Opts.AutoNumber {
		t.Errorf("Expected AutoNumber to be true")
	}

	expectedStmts := 24
	if len(diagram.Stmts) != expectedStmts {
		t.Errorf("Expected %d messages, got %d", expectedStmts, len(diagram.Stmts))
	}
	t.Run("find lifeline declaration for db", func(t *testing.T) {
		db, ok := findLifeline(diagram, "db")
		if !ok {
			t.Errorf("Expected lifelines 'db' to exist")
		} else if db.Alias != "Server Database" {
			t.Errorf("Expected db name 'Server Database', got '%s'", db.Alias)
		}
	})

	t.Run("find lifeline declaration for u", func(t *testing.T) {
		if u, ok := findLifeline(diagram, "u"); !ok {
			t.Errorf("Expected lifelines 'u' to exist")
		} else if u.Alias != "User" {
			t.Errorf("Expected 'u' name 'User', got '%s'", u.Alias)
		}
	})

	t.Run("check the first message", func(t *testing.T) {
		if m, ok := findMessage(diagram, "u", "a", "opens the website"); !ok {
			t.Errorf("First message parsed incorrectly: %+v", m)
		}
	})

	PrintDiagram(diagram)
}
