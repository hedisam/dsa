package trees

type BinaryNode struct {
	Key    int
	Left   *BinaryNode
	Right  *BinaryNode
	Parent *BinaryNode
}

func NewBinaryNode(key int) *BinaryNode {
	return &BinaryNode{Key: key}
}
