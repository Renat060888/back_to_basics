package main

import (
	"testing"

	"saod/tree"
)

func TestTreeInitialState(t *testing.T) {

	rbTree := tree.NewRBTree()

	if rbTree.GetSize() != 0 {
		t.Errorf("newly created rb-tree has non-zero size")
	}

	if !rbTree.Contains(0) {
		t.Errorf("newly created tree has some data")
	}
}

func TestTreeIsBalanced(t *testing.T) {

	rbTree := tree.NewRBTree()

	rbTree.Insert(50)

	if !rbTree.IsValid() {
		t.Errorf("rb-tree is not valid")
	}
}
