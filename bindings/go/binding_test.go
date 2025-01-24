package tree_sitter_prisma_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_prisma "github.com/victorhqc/tree-sitter-prisma/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_prisma.Language())
	if language == nil {
		t.Errorf("Error loading Prisma grammar")
	}
}
