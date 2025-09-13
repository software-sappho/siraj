# Siraj Blockchain Demo Commands

##  Run Nodes

# Node 1 (port 8080)

```bash
$ go run main.go --port=8080
```

# Node 2 (port 8081) connecting to Node 1

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

## Expose Node Publicly via ngrok (Optional)

# Expose local node on port 8080

```bash
ngrok http 8080
```

# Example usage via public URL from ngrok

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

