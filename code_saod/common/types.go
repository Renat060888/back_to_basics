package common

type INodeVisitor interface {
	VisitNode(name string, childs []string)
	// VisitNode(node *btree.BTreeNode, childs []btree.BTreeNode)
}
