package blockchain

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "sync"
    "time"
)

type VerificationRecord struct {
    UserID       string `json:"user_id"`
    DocumentHash string `json:"document_hash"`
    Status       string `json:"status"` // "verified", "pending", "rejected"
}

type Block struct {
    Index      int    `json:"index"`
    Timestamp  string `json:"timestamp"`
    PrevHash   string `json:"prev_hash"`
    MerkleRoot string `json:"merkle_root"`
    Hash       string `json:"hash"`
	Records    []VerificationRecord
}

var Blockchain []Block
var PendingRecords []VerificationRecord
var Mutex = &sync.Mutex{}

// ---------------- Hash & Block Functions ----------------

func calculateHash(block Block) string {
    record := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.PrevHash, block.MerkleRoot)
    h := sha256.Sum256([]byte(record))
    return hex.EncodeToString(h[:])
}

func IsBlockValid(newBlock, oldBlock Block) bool {
    if newBlock.Index != oldBlock.Index+1 {
        return false
    }
    if newBlock.PrevHash != oldBlock.Hash {
        return false
    }
    if calculateHash(newBlock) != newBlock.Hash {
        return false
    }
    return true
}

func GenerateBlock(prevBlock Block, merkleRoot string) Block {
    newBlock := Block{
        Index:      prevBlock.Index + 1,
        Timestamp:  time.Now().String(),
        PrevHash:   prevBlock.Hash,
        MerkleRoot: merkleRoot,
    }
    newBlock.Hash = calculateHash(newBlock)
    return newBlock
}

func AddBlock(records []VerificationRecord) Block {
    prevBlock := Blockchain[len(Blockchain)-1]
    newBlock := Block{
        Index:      prevBlock.Index + 1,
        Timestamp:  time.Now().String(),
        PrevHash:   prevBlock.Hash,
        Records:    records,                       // store records
        MerkleRoot: MerkleRoot(records),
    }
    newBlock.Hash = calculateHash(newBlock)
    Blockchain = append(Blockchain, newBlock)
    PendingRecords = []VerificationRecord{} // clear pending
    return newBlock
}


// ---------------- Genesis Block ----------------

func InitGenesisBlock() {
    genesis := Block{
        Index:      0,
        Timestamp:  time.Now().String(),
        PrevHash:   "",
        MerkleRoot: "",
        Records:    []VerificationRecord{}, // empty slice
        Hash:       "",                     // will be calculated next
    }
    genesis.Hash = calculateHash(genesis)
    Blockchain = append(Blockchain, genesis)
}
