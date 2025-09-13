package network

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"

    "siraj/blockchain"
)

var Peers []string

func BroadcastBlock(block blockchain.Block) {
    jsonBlock, _ := json.Marshal(block)
    for _, peer := range Peers {
        http.Post(fmt.Sprintf("http://%s/blocks", peer), "application/json", bytes.NewBuffer(jsonBlock))
    }
}
