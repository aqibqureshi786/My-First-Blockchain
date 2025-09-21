package main

import (
	"fmt"
	"strings"
	"time"
)

// MineBlock implements the Proof of Work algorithm
func (bc *Blockchain) MineBlock(block *Block, difficulty int) {
	target := strings.Repeat("0", difficulty)

	fmt.Printf("Mining block %d...\n", block.Index)
	startTime := time.Now()

	for {
		hash := block.calculateHash()

		if strings.HasPrefix(hash, target) {
			block.Hash = hash
			endTime := time.Now()
			duration := endTime.Sub(startTime)

			fmt.Printf("Block #%d mined successfully!\n", block.Index)
			fmt.Printf("Hash: %s\n", hash)
			fmt.Printf("Nonce: %d\n", block.Nonce)
			fmt.Printf("Mining time: %v\n", duration)
			break
		}

		block.Nonce++
	}
}

// isValidProof checks if a block satisfies the proof of work requirement
func (bc *Blockchain) isValidProof(block *Block, difficulty int) bool {
	target := strings.Repeat("0", difficulty)
	hash := block.calculateHash()
	return strings.HasPrefix(hash, target)
}

// MiningStats represents mining statistics
type MiningStats struct {
	BlockIndex   int           `json:"block_index"`
	Hash         string        `json:"hash"`
	Nonce        int           `json:"nonce"`
	Difficulty   int           `json:"difficulty"`
	MiningTime   time.Duration `json:"mining_time"`
	Transactions int           `json:"transactions"`
}

// MineBlockWithStats mines a block and returns detailed statistics
func (bc *Blockchain) MineBlockWithStats(block *Block, difficulty int) *MiningStats {
	target := strings.Repeat("0", difficulty)
	startTime := time.Now()

	for {
		hash := block.calculateHash()

		if strings.HasPrefix(hash, target) {
			block.Hash = hash
			endTime := time.Now()
			duration := endTime.Sub(startTime)

			return &MiningStats{
				BlockIndex:   block.Index,
				Hash:         hash,
				Nonce:        block.Nonce,
				Difficulty:   difficulty,
				MiningTime:   duration,
				Transactions: len(block.Data),
			}
		}

		block.Nonce++
	}
}

