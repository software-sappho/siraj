# Siraj Blockchain Demo Commands

##  Run Nodes

## Node 1 (port 8080)

```bash
$ go run main.go --port=8080
```

## Node 2 (port 8081) connecting to Node 1

```bash
go run main.go --port=8081 --peers=localhost:8080 
```
## Access User Ledger

```bash
curl http://localhost:8080/user_ledger?user_id=u123
```

## Add Verification Records

```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"user_id":"u123","document_hash":"abcd1234","status":"verified"}' \
http://localhost:8080/records
```
```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"user_id":"u1234","document_hash":"abcd5678","status":"verified"}' \
http://localhost:8080/records
```

## View Full Blockchain

```bash
curl http://localhost:8080/blocks
```

# Expose Node Publicly via ngrok (Optional)

## Expose local node on port 8080

```bash
ngrok http 8080
```

## Example usage via public URL from ngrok

```bash
curl https://abcd1234.ngrok.io/blocks
```

```bash
curl "https://abcd1234.ngrok.io/user_ledger?user_id=u123"
```

## Notes

- The blockchain is stored **in memory**, so stopping the node will reset the ledger.
- Multiple nodes sync via the `--peers` flag.
- Users can access their ledger without installing Siraj by using the ngrok public URL.

-------------------------------------------------

Guide to Building a Secure Decentralized Blockchain System

1. Define Your Blockchain Goals and Requirements
- What problem does your blockchain solve?
- What type of Blockchain? (Public, Private, Consortium)
- Concensus mechanism?(Proof of Work, Proof of Stake, PBFT or custom)
- Transaction throughput and latency targets
- Security and privacy requirements
- Integration with AI validation layers - how tightly coupled?

2. Core Blockchain Components
a) Data Structures
- Block (Index, Timestamp, Transactions, Previous Hash, Nonce etc.)
- Transaction (Sender, Receiver, Amount, Signature, Metadata)
- Blockchain (Linked list or DAG of blocks)
- State machine / ledger to track balances or asset states

b) Cryptography
- Digital signatures (e.g. ECDSA, Ed25519) for transaction authentication
- Hash functions(SHA-256, Keccak-256) for block and transaction hashing
- Public/ private key infrastucture(PKI) for identity management
- Merkle trees for efficient transaction verification

c) Consensus Algorithm
- Design or select consensus protocol appropriate for your trust model
- Implement leader election, block proposal, voting or mining
- Handle forks and chain reorganizations
- Include incentive and punishment mechanisms (if applicable)

3. Networking Layer
- Peer-to-peer (P2P) network protocol
- Node discovery and bootstrapping
- Message propagation( transactions, blocks, status updates)
- Handling network partitions and latency
- Secure communications channels (TLS or encrypted transport)

4. Node Architecture
- Full nodes: store full blockchain and validate all blocks/transactions
- Light nodes: store minimal data for efficient use cases
- Validator/miner modes: participate actively in consensus
- APIs for node interaction (RPC, REST, gRPC)

5. Transaction and Block Validation
- Verify transaction signatures and balances
- Enforce rules (double-spend protection, nonce increments)
- Validate block contents and proof of consensus
- Reject invalid or malicious data

6. Persistence and Storage
- Choose a database or storage engine (LevelDB, RocksDB, BoltDB)
- Efficient indexing of blocks and transactions
- State snapshots and pruning strategies
- Backup and recovery mechanisms

7. Security Hardening
- Input validation and sanitization to prevent injection attacks
- Secure key management and storage (hardware wallets or HSM integration)
- Rate limiting and DoS protection
- Protection  against Sybil attacks and network flooding
- Encryption of network traffic and sensitive data
- Implement auditing, logging and anomaly detection

8. Integration with AI-Driven Validation Layers
- Define interfaces/APIs between AI layers and blockchain
- AI layers may preprocess or score transactions/blocks
- Use blockchain for immutable audit trails of AI decisions
- Ensure trustworthiness and transparency in AI validations

9. Concensus and Security Testing
- Unit and integration tests for blockchain functions
- Network simulation to test consensus under partition/failures
- Security audits and formal verification(where possible)
- Penetration testing for network and node vulnerabilities

10. Deployment and Maintenance
- Containerize nodes for easier deployment(Docker, Kubernetes)
- Set up monitoring, alerting and logging (Prometheus, Grafana)
- Implement automatic updates and patch management
- Plan for scaling network and adding nodes dynamically
- Establish governance and upgrade mechanisms

11. Documentation and Developer Tools
- Comprehensive technical documentation
SDKs and libraries for interacting with your blockchain
- CLI tools for node operators
- Developer environment setup and tutorials

XXXX SUGGESTED DEVELOPMENT ROADMAP XXX

1. Prototype blockchain core - blocks, chain, validation

2. Implement networking - P2P connections and messaging

3. Add consensus algorithm - simple PoW or PoS for testing

4. Build transaction pool and validation logic

5. Integrate AI-driven validation layer - initially as a separate microservice

6. Add persistence layer

7. Implement node APIs and CLI tools

8. Conduct security hardening and testing

9. Deploy multi-node test network

10. Iterate, optimize and prepare for mainnet launch.