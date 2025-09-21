package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var blockchain *Blockchain

func main() {
	// Initialize blockchain
	blockchain = NewBlockchain()
	fmt.Println("Aqib Mehmood Qureshi Blockchain initialized with genesis block")

	// Setup routes
	router := mux.NewRouter()

	// API endpoints
	router.HandleFunc("/api/blockchain", getBlockchain).Methods("GET")
	router.HandleFunc("/api/blockchain/info", getBlockchainInfo).Methods("GET")
	router.HandleFunc("/api/blocks", getBlocks).Methods("GET")
	router.HandleFunc("/api/blocks/{index}", getBlock).Methods("GET")
	router.HandleFunc("/api/mine", mineBlock).Methods("POST")
	router.HandleFunc("/api/search", searchData).Methods("GET")
	router.HandleFunc("/api/validate", validateChain).Methods("GET")

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // React app
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(router)

	fmt.Println("Server starting on port 8080...")
	fmt.Println("API endpoints available:")
	fmt.Println("  GET  /api/blockchain - Get full blockchain")
	fmt.Println("  GET  /api/blockchain/info - Get blockchain info")
	fmt.Println("  GET  /api/blocks - Get all blocks")
	fmt.Println("  GET  /api/blocks/{index} - Get specific block")
	fmt.Println("  POST /api/mine - Mine a new block")
	fmt.Println("  GET  /api/search?q={query} - Search for data")
	fmt.Println("  GET  /api/validate - Validate blockchain")

	log.Fatal(http.ListenAndServe(":8080", handler))
}

// getBlockchain returns the entire blockchain
func getBlockchain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blockchain)
}

// getBlockchainInfo returns blockchain statistics
func getBlockchainInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	info := blockchain.GetBlockchainInfo()
	json.NewEncoder(w).Encode(info)
}

// getBlocks returns all blocks
func getBlocks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blockchain.Blocks)
}

// getBlock returns a specific block by index
func getBlock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	index, err := strconv.Atoi(vars["index"])
	if err != nil {
		http.Error(w, "Invalid block index", http.StatusBadRequest)
		return
	}

	block := blockchain.GetBlockByIndex(index)
	if block == nil {
		http.Error(w, "Block not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(block)
}

// MineRequest represents a mining request
type MineRequest struct {
	Data []string `json:"data"`
}

// mineBlock mines a new block with the provided data
func mineBlock(w http.ResponseWriter, r *http.Request) {
	var req MineRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if len(req.Data) == 0 {
		http.Error(w, "No data provided for mining", http.StatusBadRequest)
		return
	}

	// Create new block
	prevBlock := blockchain.GetLatestBlock()
	newIndex := len(blockchain.Blocks)
	newBlock := NewBlock(newIndex, req.Data, prevBlock.Hash)

	// Mine the block with stats
	stats := blockchain.MineBlockWithStats(newBlock, 4)

	// Add to blockchain
	blockchain.Blocks = append(blockchain.Blocks, newBlock)

	// Return mining stats and block info
	response := map[string]interface{}{
		"success": true,
		"block":   newBlock,
		"stats":   stats,
		"message": fmt.Sprintf("Block #%d mined successfully", newBlock.Index),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// searchData searches for data in the blockchain
func searchData(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Search query is required", http.StatusBadRequest)
		return
	}

	results := blockchain.SearchData(query)

	response := map[string]interface{}{
		"query":   query,
		"results": results,
		"count":   len(results),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// validateChain validates the entire blockchain
func validateChain(w http.ResponseWriter, r *http.Request) {
	isValid := blockchain.IsChainValid()

	response := map[string]interface{}{
		"is_valid": isValid,
		"message":  getValidationMessage(isValid),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getValidationMessage(isValid bool) string {
	if isValid {
		return "Blockchain is valid and secure"
	}
	return "Blockchain validation failed - integrity compromised"
}
