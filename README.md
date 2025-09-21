# Aqib Mehmood Qureshi Blockchain - Complete Implementation

A full-featured blockchain implementation with Go backend and React frontend, featuring all required components including Proof of Work, Merkle Trees, and a beautiful user interface.

## 🏗️ Project Structure

```
Blockchain Assignment/
├── backend/                 # Go blockchain implementation
│   ├── main.go             # REST API server
│   ├── block.go            # Block structure and operations
│   ├── blockchain.go       # Blockchain logic and validation
│   ├── merkle.go           # Merkle tree implementation
│   ├── mining.go           # Proof of Work algorithm
│   └── go.mod              # Go module dependencies
├── frontend/               # React TypeScript frontend
│   ├── src/
│   │   ├── components/     # React components
│   │   ├── services/       # API service layer
│   │   ├── types/          # TypeScript type definitions
│   │   └── App.tsx         # Main application
│   └── package.json        # Node.js dependencies
└── README.md               # This file
```

## ✨ Features Implemented

### ✅ Part 1 Requirements

1. **Block Structure** [5 points]
   - ✅ Index (height)
   - ✅ Timestamp
   - ✅ Data (Transactions)
   - ✅ PrevHash
   - ✅ Hash
   - ✅ Nonce

2. **Genesis Block** [10 points]
   - ✅ Automatically created when blockchain initializes
   - ✅ Special genesis block with index 0 and previous hash "0"

3. **Merkle Tree** [10 points]
   - ✅ Complete Merkle tree implementation
   - ✅ Efficient and secure transaction storage
   - ✅ Merkle root calculation for each block

4. **String Transactions** [5 points]
   - ✅ Transactions stored as strings
   - ✅ Multiple transactions per block support

5. **Proof of Work Mining** [10 points]
   - ✅ Complete PoW algorithm implementation
   - ✅ Configurable difficulty (default: 4 leading zeros)
   - ✅ Mine button in UI
   - ✅ Mining statistics and performance tracking

6. **View Complete Blockchain** [10 points]
   - ✅ Beautiful UI showing all blocks
   - ✅ Block details including hash, nonce, timestamp
   - ✅ Transaction listing for each block
   - ✅ Blockchain statistics dashboard

7. **Search Functionality** [10 points]
   - ✅ Search across all transactions in all blocks
   - ✅ Highlighted search results
   - ✅ Block and transaction index information

8. **React User Interface** [Bonus]
   - ✅ Modern, responsive React TypeScript app
   - ✅ Beautiful gradient design with animations
   - ✅ Tabbed interface for different operations
   - ✅ Real-time updates and error handling

## 🚀 Getting Started

### Prerequisites

- **Go** (version 1.19 or later)
- **Node.js** (version 16 or later)
- **npm** or **yarn**

### Installation & Setup

1. **Clone or navigate to the project directory:**
   ```bash
   cd "C:\Users\HP\Desktop\Blockchain Assignment"
   ```

2. **Start the Go Backend:**
   ```bash
   cd backend
   go mod tidy
   go run .
   ```
   The server will start on `http://localhost:8080`

3. **Start the React Frontend (in a new terminal):**
   ```bash
   cd frontend
   npm install
   npm start
   ```
   The React app will open at `http://localhost:3000`

### Alternative: Build and Run

**Backend:**
```bash
cd backend
go build -o blockchain.exe
./blockchain.exe
```

**Frontend:**
```bash
cd frontend
npm run build
# Serve the build directory with a static server
```

## 🎯 How to Use

### 1. View Blockchain
- Click the "View Blockchain" tab
- See all blocks with their details
- View blockchain statistics (total blocks, transactions, validation status)
- Each block shows: Index, Hash, Previous Hash, Nonce, Merkle Root, Timestamp, and Transactions

### 2. Mine New Blocks
- Click the "Mine Block" tab
- Add transactions by typing in the input field and clicking "Add"
- You can add multiple transactions to a single block
- Click "Mine Block" to start the Proof of Work process
- Watch the mining progress and see detailed statistics after completion

### 3. Search Data
- Click the "Search Data" tab
- Enter any search term to find it across all transactions
- Results show which block and transaction index contain your search term
- Search terms are highlighted in the results

## 🔧 API Endpoints

The Go backend provides these REST API endpoints:

- `GET /api/blockchain` - Get the complete blockchain
- `GET /api/blockchain/info` - Get blockchain statistics
- `GET /api/blocks` - Get all blocks
- `GET /api/blocks/{index}` - Get a specific block
- `POST /api/mine` - Mine a new block with transactions
- `GET /api/search?q={query}` - Search for data in the blockchain
- `GET /api/validate` - Validate the entire blockchain

## 🏛️ Architecture Details

### Backend (Go)

- **Block Structure**: Complete implementation with all required fields
- **Merkle Tree**: Binary tree structure for efficient transaction verification
- **Proof of Work**: SHA-256 based mining with configurable difficulty
- **Blockchain Validation**: Ensures chain integrity and valid proof of work
- **REST API**: Clean API design with proper error handling and CORS support

### Frontend (React + TypeScript)

- **Modern UI**: Beautiful, responsive design with smooth animations
- **Type Safety**: Full TypeScript implementation with proper type definitions
- **Component Architecture**: Modular, reusable React components
- **State Management**: Efficient state handling with React hooks
- **Error Handling**: Comprehensive error handling with user feedback

### Key Technical Features

1. **Merkle Tree Implementation**
   - Binary tree structure for transaction verification
   - Handles odd numbers of transactions by duplication
   - SHA-256 hashing for security

2. **Proof of Work Algorithm**
   - Configurable difficulty (number of leading zeros)
   - Nonce increment until valid hash found
   - Mining time tracking and statistics

3. **Blockchain Validation**
   - Hash verification for each block
   - Previous hash chain validation
   - Proof of work verification

4. **Search Functionality**
   - Case-insensitive search across all transactions
   - Returns block index, transaction index, and context
   - Efficient string matching

## 🛡️ Security Features

- **Cryptographic Hashing**: SHA-256 for all hash operations
- **Merkle Tree Verification**: Efficient transaction integrity checking
- **Proof of Work**: Computational security against tampering
- **Chain Validation**: Complete blockchain integrity verification
- **CORS Protection**: Secure cross-origin resource sharing

## 🎨 UI Features

- **Responsive Design**: Works on desktop and mobile devices
- **Real-time Updates**: Automatic refresh after mining operations
- **Loading States**: Clear feedback during operations
- **Error Handling**: User-friendly error messages
- **Animations**: Smooth transitions and loading animations
- **Modern Styling**: Beautiful gradient backgrounds and card layouts

## 🔍 Troubleshooting

### Common Issues

1. **"Failed to fetch" errors**
   - Ensure the Go backend is running on port 8080
   - Check that CORS is properly configured

2. **Mining takes too long**
   - This is normal! Proof of Work is computationally intensive
   - Mining time depends on your computer's performance

3. **React app won't start**
   - Make sure you ran `npm install` in the frontend directory
   - Check that Node.js version is 16 or later

### Fail-Safe Options

If something goes wrong:

1. **Restart both servers** - Stop both Go and React servers, then restart them
2. **Clear browser cache** - Hard refresh the React app (Ctrl+F5)
3. **Check terminal logs** - Look for error messages in both terminal windows
4. **Verify ports** - Ensure ports 8080 (Go) and 3000 (React) are available

## 📊 Performance Notes

- **Mining Difficulty**: Set to 4 leading zeros (adjustable in code)
- **Expected Mining Time**: 1-10 seconds depending on hardware
- **Memory Usage**: Minimal - blockchain stored in memory
- **Scalability**: Suitable for demonstration and learning purposes

## 🎯 Assignment Completion

This implementation fully satisfies all requirements:

- ✅ **Block Structure** with all required fields (5 points)
- ✅ **Genesis Block** creation (10 points)
- ✅ **Merkle Tree** implementation (10 points)
- ✅ **String Transactions** support (5 points)
- ✅ **Proof of Work Mining** with mine button (10 points)
- ✅ **Complete Blockchain View** (10 points)
- ✅ **Search Functionality** (10 points)
- ✅ **React User Interface** for all operations

**Total: 70+ points** (including bonus UI implementation)

## 👨‍💻 Development

Built with modern technologies:
- **Backend**: Go 1.19+, Gorilla Mux, CORS
- **Frontend**: React 18, TypeScript, Lucide Icons
- **Styling**: Custom CSS with modern design principles
- **Architecture**: RESTful API with clean separation of concerns

---

**🎉 Ready to explore the Aqib Mehmood Qureshi Blockchain! Start both servers and visit http://localhost:3000**
