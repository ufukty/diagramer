package sequence

import (
	"bytes"
	"os"
	"testing"
)

func TestRender(t *testing.T) {
	src, err := os.Open("testdata/1.txt")
	if err != nil {
		t.Fatalf("arrange, open file: %v", err)
	}
	defer src.Close() //nolint:errcheck

	dst := bytes.NewBufferString("")
	err = Render(dst, src)
	if err != nil {
		t.Fatalf("act, Render: %v", err)
	}

	panic("add assertions") // TODO:
}
