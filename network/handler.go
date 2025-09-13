package network

import (
    "encoding/json"
    "io/ioutil"
    "net/http"

    "siraj/blockchain"
)

// GET /blocks returns the full blockchain
func HandleGetBlockchain(w http.ResponseWriter, r *http.Request) {
    blockchain.Mutex.Lock()
    defer blockchain.Mutex.Unlock()
    json.NewEncoder(w).Encode(blockchain.Blockchain)
}

// POST /blocks receives a new block from peer
func HandleReceiveBlock(w http.ResponseWriter, r *http.Request) {
    var newBlock blockchain.Block
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &newBlock)

    blockchain.Mutex.Lock()
    defer blockchain.Mutex.Unlock()
    if blockchain.IsBlockValid(newBlock, blockchain.Blockchain[len(blockchain.Blockchain)-1]) {
        blockchain.Blockchain = append(blockchain.Blockchain, newBlock)
    }
}

// POST /records adds a new verification record
func HandleAddRecord(w http.ResponseWriter, r *http.Request) {
    var record blockchain.VerificationRecord
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &record)

    blockchain.Mutex.Lock()
    blockchain.PendingRecords = append(blockchain.PendingRecords, record)
    blockchain.Mutex.Unlock()

    if len(blockchain.PendingRecords) >= 2 {
    newBlock := blockchain.AddBlock(blockchain.PendingRecords)
    BroadcastBlock(newBlock)
}


    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Record added to Siraj pending list"))
}

func HandleUserLedger(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("user_id")
    if userID == "" {
        http.Error(w, "user_id required", http.StatusBadRequest)
        return
    }

    var userRecords []blockchain.VerificationRecord

    blockchain.Mutex.Lock()
    defer blockchain.Mutex.Unlock()

    // 1. Check pending records
    for _, rec := range blockchain.PendingRecords {
        if rec.UserID == userID {
            userRecords = append(userRecords, rec)
        }
    }

    // 2. Check records in blocks
    for _, block := range blockchain.Blockchain {
        for _, rec := range block.Records {
            if rec.UserID == userID {
                userRecords = append(userRecords, rec)
            }
        }
    }

    json.NewEncoder(w).Encode(userRecords)
}
