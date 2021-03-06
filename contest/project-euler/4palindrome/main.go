package main
import (
	"fmt"


	"errors"
	"strings"
)

func main() {
	var t int
	fmt.Scan(&t)
	//	q := make(chan bool)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		process(n)
	}
}

func process(t int) {
	//	q := make(chan bool)

	bst := BST{}
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			num := i * j
				if num <= t && num >100000 && isPalindrom(num) {
					bst.Insert(int64(num))
				}
		}

	}

	max, _ := bst.Max();
	fmt.Println(max)
}

func isPalindrom(n int) bool {
	return n == reverse(n)
}

func reverse(n int) (rev int) {

	for n != 0 {
		rev = rev * 10 + n % 10
		n /= 10
	}
	return rev
}



type Node struct {
	Value int64
	Left  *Node
	Right *Node
}

type BST struct {
	root *Node
}


func (b *BST) Insert(value int64) {

	b.root = insert(b.root, value)
}

func (b *BST) Min() (int64, error) {
	if b.root != nil {
		return b.root.FindMin().Value, nil
	}
	return 0, errors.New("BST is empty")
}

func (b *BST) Max() (int64, error) {
	if b.root != nil {
		return b.root.FindMax().Value, nil
	}
	return 0, errors.New("BST is empty")
}

func (b *BST) Reverse(){
	b.root.Reverse()
}

func (b *BST) Traverse() {
	b.root.Traverse(0)
}

func (b *BST) Delete(key int64) {
	if b.root != nil {
		b.root = deleteNode(b.root, key)
	}
}

func (n *Node) Traverse(level int) {
	if n.Right != nil { n.Right.Traverse(level+1) }
	fmt.Printf("> %s%d\n", strings.Repeat("\t", level), n.Value)
	if n.Left != nil { n.Left.Traverse(level+1) }
}
func (n *Node) Reverse() {
	n.Right, n.Left = n.Left, n.Right
	if n.Right != nil { n.Right.Reverse() }
	if n.Left != nil { n.Left.Reverse() }
}

func (n *Node) FindMin() *Node {
//	fmt.Println("min", n.Value)
	if n.Left != nil {
		return n.Left.FindMin()
	}
	return n
}

func (n *Node) FindMax() *Node {
//	fmt.Println("max", n.Value)
	if n.Right != nil {
		return n.Right.FindMax()
	}
	return n
}

func insert(n *Node, value int64) *Node {
	if n == nil {
		n = &Node{Value:value}
	} else if value > n.Value {
		if n.Right == nil {
			n.Right = &Node{Value:value}
		}
		n.Right = insert(n.Right, value)
	} else if value < n.Value  {
		if n.Left == nil {
			n.Left = &Node{Value:value}
		}
		n.Left = insert(n.Left, value)
	}
	return n
}

func deleteMax(n *Node) *Node {
	if n.Right == nil {
		return n.Left
	}
	n.Right = deleteMax(n.Right)
	return n
}

func deleteMin(n *Node) *Node {
	if n.Left == nil {
		return n.Right
	}
	n.Left = deleteMin(n.Left)
	return n
}

func deleteNode(n *Node, key int64) *Node {
	//	fmt.Printf("%+v\n", n)
	if key < n.Value  {
		n.Left = deleteNode(n.Left, key)
	} else if key > n.Value  {
		n.Right = deleteNode(n.Right, key)
	} else {
		// Remove this node
		if n.Left == nil && n.Right == nil {
			return nil
		}
		if n.Left == nil {
			return n.Right
		}
		if n.Right == nil {
			return n.Left
		}
		//		n.Traverse(0)
		t := n
		n = t.Right.FindMin()
		n.Right = deleteMin(t.Right)
		n.Left = t.Left

		//		fmt.Println("right")
		//		n.Right.Traverse(1)
		//		fmt.Println("left")
		//		n.Left.Traverse(1)
		//		fmt.Printf("%+v, left %+v, right %+v", n, n.Left, n.Right)
	}

	return n
}


