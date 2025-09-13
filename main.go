package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"siraj/blockchain"
	"siraj/network"
)

func main() {
	port := flag.String("port", "8080", "server port")
	peersList := flag.String("peers", "", "comma separated list of peers")
	flag.Parse()

	if *peersList != "" {
		network.Peers = strings.Split(*peersList, ",")
	}

	blockchain.InitGenesisBlock()

	http.HandleFunc("/blocks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			network.HandleGetBlockchain(w, r)
		} else if r.Method == "POST" {
			network.HandleReceiveBlock(w, r)
		}
	})
	http.HandleFunc("/records", network.HandleAddRecord)
	http.HandleFunc("/user_ledger", network.HandleUserLedger)

	fmt.Println("Siraj node running on port", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
