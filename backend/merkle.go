package main

import (
	"crypto/sha256"
	"encoding/hex"
)

// MerkleNode represents a node in the Merkle tree
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  string
	Hash  string
}

// NewMerkleNode creates a new Merkle node
func NewMerkleNode(left, right *MerkleNode, data string) *MerkleNode {
	node := &MerkleNode{
		Left:  left,
		Right: right,
		Data:  data,
	}

	if left == nil && right == nil {
		// Leaf node
		hash := sha256.Sum256([]byte(data))
		node.Hash = hex.EncodeToString(hash[:])
	} else {
		// Internal node
		prevHashes := left.Hash + right.Hash
		hash := sha256.Sum256([]byte(prevHashes))
		node.Hash = hex.EncodeToString(hash[:])
	}

	return node
}

// MerkleTree represents the Merkle tree structure
type MerkleTree struct {
	Root *MerkleNode
}

// NewMerkleTree creates a new Merkle tree from a list of data
func NewMerkleTree(data []string) *MerkleTree {
	if len(data) == 0 {
		return &MerkleTree{Root: NewMerkleNode(nil, nil, "")}
	}

	var nodes []*MerkleNode

	// Create leaf nodes
	for _, datum := range data {
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, node)
	}

	// If odd number of nodes, duplicate the last one
	if len(nodes)%2 != 0 {
		nodes = append(nodes, nodes[len(nodes)-1])
	}

	// Build the tree bottom-up
	for len(nodes) > 1 {
		var newLevel []*MerkleNode

		for i := 0; i < len(nodes); i += 2 {
			node := NewMerkleNode(nodes[i], nodes[i+1], "")
			newLevel = append(newLevel, node)
		}

		// If odd number of nodes in this level, duplicate the last one
		if len(newLevel)%2 != 0 && len(newLevel) > 1 {
			newLevel = append(newLevel, newLevel[len(newLevel)-1])
		}

		nodes = newLevel
	}

	return &MerkleTree{Root: nodes[0]}
}

// calculateMerkleRoot calculates the Merkle root for the block's transactions
func (b *Block) calculateMerkleRoot() string {
	if len(b.Data) == 0 {
		return ""
	}

	tree := NewMerkleTree(b.Data)
	return tree.Root.Hash
}
