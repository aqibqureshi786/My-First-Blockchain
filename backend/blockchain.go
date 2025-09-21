package main

import (
	"fmt"
	"strings"
)

// Blockchain represents the entire blockchain
type Blockchain struct {
	Blocks []*Block `json:"blocks"`
}

// NewBlockchain creates a new blockchain with genesis block
func NewBlockchain() *Blockchain {
	blockchain := &Blockchain{}
	blockchain.CreateGenesisBlock()
	return blockchain
}

// CreateGenesisBlock creates the first block in the blockchain
func (bc *Blockchain) CreateGenesisBlock() {
	genesisData := []string{"Genesis Block - The beginning of Aqib Mehmood Qureshi Blockchain"}
	genesisBlock := NewBlock(0, genesisData, "0")
	genesisBlock.Hash = genesisBlock.calculateHash()
	bc.Blocks = append(bc.Blocks, genesisBlock)
}

// GetLatestBlock returns the last block in the blockchain
func (bc *Blockchain) GetLatestBlock() *Block {
	if len(bc.Blocks) == 0 {
		return nil
	}
	return bc.Blocks[len(bc.Blocks)-1]
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data []string) {
	prevBlock := bc.GetLatestBlock()
	newIndex := len(bc.Blocks)
	newBlock := NewBlock(newIndex, data, prevBlock.Hash)

	// Mine the block
	bc.MineBlock(newBlock, 4) // difficulty of 4

	bc.Blocks = append(bc.Blocks, newBlock)
}

// IsChainValid validates the entire blockchain
func (bc *Blockchain) IsChainValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		prevBlock := bc.Blocks[i-1]

		// Check if current block's hash is valid
		if currentBlock.Hash != currentBlock.calculateHash() {
			fmt.Printf("Invalid hash at block %d\n", i)
			return false
		}

		// Check if current block points to previous block
		if currentBlock.PrevHash != prevBlock.Hash {
			fmt.Printf("Invalid previous hash at block %d\n", i)
			return false
		}

		// Check if block satisfies proof of work
		if !bc.isValidProof(currentBlock, 4) {
			fmt.Printf("Invalid proof of work at block %d\n", i)
			return false
		}
	}
	return true
}

// SearchData searches for data across all blocks in the blockchain
func (bc *Blockchain) SearchData(query string) []*SearchResult {
	var results []*SearchResult
	query = strings.ToLower(query)

	for _, block := range bc.Blocks {
		for i, data := range block.Data {
			if strings.Contains(strings.ToLower(data), query) {
				result := &SearchResult{
					BlockIndex:       block.Index,
					TransactionIndex: i,
					Data:             data,
					BlockHash:        block.Hash,
					Timestamp:        block.Timestamp,
				}
				results = append(results, result)
			}
		}
	}

	return results
}

// SearchResult represents a search result
type SearchResult struct {
	BlockIndex       int         `json:"block_index"`
	TransactionIndex int         `json:"transaction_index"`
	Data             string      `json:"data"`
	BlockHash        string      `json:"block_hash"`
	Timestamp        interface{} `json:"timestamp"`
}

// GetBlockByIndex returns a block by its index
func (bc *Blockchain) GetBlockByIndex(index int) *Block {
	if index < 0 || index >= len(bc.Blocks) {
		return nil
	}
	return bc.Blocks[index]
}

// GetBlockByHash returns a block by its hash
func (bc *Blockchain) GetBlockByHash(hash string) *Block {
	for _, block := range bc.Blocks {
		if block.Hash == hash {
			return block
		}
	}
	return nil
}

// GetBlockchainInfo returns general information about the blockchain
func (bc *Blockchain) GetBlockchainInfo() map[string]interface{} {
	totalTransactions := 0
	for _, block := range bc.Blocks {
		totalTransactions += len(block.Data)
	}

	return map[string]interface{}{
		"total_blocks":       len(bc.Blocks),
		"total_transactions": totalTransactions,
		"is_valid":           bc.IsChainValid(),
		"latest_block_hash":  bc.GetLatestBlock().Hash,
	}
}
