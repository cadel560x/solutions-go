package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	// Tree data structure where attempts are stored
	guessingTree := &tree{}

	// s1 := rand.NewSource(time.Now().UnixNano())
	// randomNumber := rand.New(s1)

	// Debug
	randomNumber := uint16(rand.Intn(9999) + 1)
	var triedNumber uint16
	var found bool
	var tryCounter, previousTriedNumber, foundValue uint16

	// Title
	fmt.Println("Guessing Game")
	fmt.Println("-------------\n")

	// Instructions
	fmt.Println("Guess a number positive integer between 1 and 10000 inclusive.")
	fmt.Println("Press '0' to exit.\n")

	// Debug
	fmt.Printf("Debug: randomNumber: %d\n\n", randomNumber)

	// Let's make the 'randomNumber' the root of the 'guessingTree'
	guessingTree.insert(randomNumber)

	// Debug
	// fmt.Println("Debug: tree: traverse:")
	// guessingTree.traverse(guessingTree.Root, func(n *node) { fmt.Printf("Value: %d | ", n.Value) })
	// fmt.Println()


	// Capture the first try
	fmt.Print("Guess a number: ")
	fmt.Scan(&triedNumber)

	// Main game loop
	for triedNumber != 0 {
		// Debug
		fmt.Printf("Debug: triedNumber: %d\n\n", triedNumber)

		foundValue, found = guessingTree.find(triedNumber)
		// Debug
		fmt.Println("Debug: tree: find: ", found, "\n")

		// 'randomNumber' was guessed
		if foundValue == randomNumber {
			// Count this try
			tryCounter++
			// Debug
			fmt.Println("Debug: for loop: tryCounter: ", tryCounter, "\n")

			// Force loop exit
			// triedNumber = 0
			break;

		// 'triedNumber' is found in 'guessingTree'
		} else if foundValue != randomNumber && found {
			// Debug
			fmt.Println("Debug: for loop: previousTriedNumber: ", previousTriedNumber, "\n")

			// Avoid counting the same number multiple times consecutively
			if previousTriedNumber != triedNumber {
				tryCounter++
				// Debug
				fmt.Println("Debug: for loop: tryCounter: ", tryCounter, "\n")
			}

			// Save 'triedNumber' value for next iteration
			previousTriedNumber = triedNumber
			// Debug
			fmt.Println("Debug: for loop: previousTriedNumber: ", previousTriedNumber)
			fmt.Println("Debug: for loop: triedNumber: ", triedNumber)
			fmt.Println("Debug: for loop: tryCounter: ", tryCounter, "\n")

		// 'triedNumber' is not in 'guessingTree'
		} else {
			// Add the new 'triedNumber' into the 'guessingTree'
			guessingTree.insert(triedNumber)
			// Debug
			fmt.Println("Debug: tree: traverse:")
			guessingTree.traverse(guessingTree.Root, func(n *node) { fmt.Printf("Value: %d | ", n.Value) })
			fmt.Println()

			// Save 'triedNumber' value for next iteration
			previousTriedNumber = triedNumber
			tryCounter++
			// Debug
			fmt.Println("Debug: for loop: previousTriedNumber: ", previousTriedNumber)
			fmt.Println("Debug: for loop: triedNumber: ", triedNumber)
			fmt.Println("Debug: for loop: tryCounter: ", tryCounter, "\n")

		} // if - else if - else

		// Subsequent capture
		fmt.Print("Guess a number: ")
		fmt.Scan(&triedNumber)

	} // for - main game loop

} // main

// I searched for a Go implementation of a binary tree in Go's API and didn't find one.
// Tree data structure taken and adapted from https://appliedgo.net/bintree/

// Node
type node struct {
	Value uint16
	Left  *node
	Right *node
} // node

func (n *node) insert(value uint16) error {
	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}

	switch {
	case value == n.Value:
		return nil
	case value < n.Value:
		if n.Left == nil {
			n.Left = &node{Value: value}
			return nil
		}
		return n.Left.insert(value)
	case value > n.Value:
		if n.Right == nil {
			n.Right = &node{Value: value}
			return nil
		}
		return n.Right.insert(value)
	}
	return nil
} // node::insert

func (n *node) find(s uint16) (uint16, bool) {

	if n == nil {
		// Return value of '0' means 'not found'
		return 0, false
	}

	switch {
	case s == n.Value:
		// Debug
		fmt.Println("Debug: node: find: s: ", s)
		return n.Value, true
	case s < n.Value:
		return n.Left.find(s)
	default:
		return n.Right.find(s)
	}
} // node::find

func (n *node) findMax(parent *node) (*node, *node) {
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(n)
} // node::findMax

// replaceNode replaces the parent’s child pointer to n with a pointer to the replacement node. parent must not be nil.
func (n *node) replaceNode(parent, replacement *node) error {
	if n == nil {
		return errors.New("replaceNode() not allowed on a nil node")
	}

	if n == parent.Left {
		parent.Left = replacement
		return nil
	}
	parent.Right = replacement
	return nil
} // node::replaceNode

// Delete removes an element from the tree. It is an error to try deleting an element that does not exist. In order to remove an element properly, Delete needs to know the node’s parent node. parent must not be nil.
func (n *node) delete(s uint16, parent *node) error {
	if n == nil {
		return errors.New("Value to be deleted does not exist in the tree")
	}
	// Search the node to be deleted.
	switch {
	case s < n.Value:
		return n.Left.delete(s, n)
	case s > n.Value:
		return n.Right.delete(s, n)
	default:
		// We found the node to be deleted. If the node has no children, simply remove it from its parent.
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
			return nil
		}
		// If the node has one child: Replace the node with its child.
		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}
		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return nil
		}
		// If the node has two children: Find the maximum element in the left subtree…
		replacement, replParent := n.Left.findMax(n)
		// …and replace the node’s value and data with the replacement’s value and data.
		n.Value = replacement.Value
		// Then remove the replacement node.
		return replacement.delete(replacement.Value, replParent)
	}
} // node::delete

// Tree
type tree struct {
	Root *node
}

func (t *tree) insert(value uint16) error {
	// If the tree is empty, create a new node,…
	if t.Root == nil {
		t.Root = &node{Value: value}
		return nil
	}
	// …else call Node.Insert.
	return t.Root.insert(value)
} // tree::insert

// Find calls Node.Find unless the root node is nil
func (t *tree) find(s uint16) (uint16, bool) {
	if t.Root == nil {
		// Return value of '0' means 'not found'
		return 0, false
	}
	return t.Root.find(s)
} // tree::find

// Delete has one special case: the empty tree. (And deleting from an empty tree is an error.) In all other cases, it calls Node.Delete.
func (t *tree) delete(s uint16) error {

	if t.Root == nil {
		return errors.New("Cannot delete from an empty tree")
	}
	// CallNode.Delete. Passing a “fake” parent node here almost avoids having to treat the root node as a special case, with one exception.
	fakeParent := &node{Right: t.Root}
	err := t.Root.delete(s, fakeParent)
	if err != nil {
		return err
	}
	// If the root node is the only node in the tree, and if it is deleted, then it only got removed from fakeParent. t.Root still points to the old node. We rectify this by setting t.Root to nil.
	if fakeParent.Right == nil {
		t.Root = nil
	}
	return nil
} // tree::delete

// Traverse is a simple method that traverses the tree in left-to-right order (which, by pure incidence ;-), is the same as traversing from smallest to largest value) and calls a custom function on each node.
func (t *tree) traverse(n *node, f func(*node)) {
	if n == nil {
		return
	}
	t.traverse(n.Left, f)
	f(n)
	t.traverse(n.Right, f)
} // tree::traverse
