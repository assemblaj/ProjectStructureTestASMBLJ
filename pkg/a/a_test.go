package a

import (
	"testing"

	c "github.com/assemblaj/ProjectStructureTestASMBLJ/pkg/a/internal/"
	b "github.com/assemblaj/ProjectStructureTestASMBLJ/pkg/b"
)

func TestPrintA(t *testing.T) {
	PrintA()
}

func TestPrintB(t *testing.T) {
	b.PrintB()
}

func TestPrintC(t *testing.T) {
	c.PrintC()
}
