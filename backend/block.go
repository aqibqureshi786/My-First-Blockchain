package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

// Block represents a single block in the blockchain
type Block struct {
	Index      int       `json:"index"`
	Timestamp  time.Time `json:"timestamp"`
	Data       []string  `json:"data"`
	PrevHash   string    `json:"prev_hash"`
	Hash       string    `json:"hash"`
	Nonce      int       `json:"nonce"`
	MerkleRoot string    `json:"merkle_root"`
}

// NewBlock creates a new block with the given data and previous hash
func NewBlock(index int, data []string, prevHash string) *Block {
	block := &Block{
		Index:     index,
		Timestamp: time.Now(),
		Data:      data,
		PrevHash:  prevHash,
		Nonce:     0,
	}

	// Calculate Merkle root for the transactions
	block.MerkleRoot = block.calculateMerkleRoot()

	return block
}

// calculateHash generates the hash for the block
func (b *Block) calculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp.String() + b.MerkleRoot + b.PrevHash + strconv.Itoa(b.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// String returns a string representation of the block
func (b *Block) String() string {
	return fmt.Sprintf("Block #%d [%s] - Hash: %s, PrevHash: %s, Nonce: %d, Data: %v",
		b.Index, b.Timestamp.Format("2006-01-02 15:04:05"), b.Hash, b.PrevHash, b.Nonce, b.Data)
}

