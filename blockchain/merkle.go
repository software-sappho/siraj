package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

func hashData(data string) []byte {
    h := sha256.Sum256([]byte(data))
    return h[:]
}

// Compute Merkle Root for a batch of verification records
func MerkleRoot(records []VerificationRecord) string {
    var leaves [][]byte
    for _, r := range records {
        combined := r.UserID + r.DocumentHash + r.Status
        leaves = append(leaves, hashData(combined))
    }

    if len(leaves) == 0 {
        return ""
    }

    for len(leaves) > 1 {
        var nextLevel [][]byte
        for i := 0; i < len(leaves); i += 2 {
            if i+1 < len(leaves) {
                combined := append(leaves[i], leaves[i+1]...)
                nextLevel = append(nextLevel, hashData(string(combined)))
            } else {
                combined := append(leaves[i], leaves[i]...)
                nextLevel = append(nextLevel, hashData(string(combined)))
            }
        }
        leaves = nextLevel
    }

    return hex.EncodeToString(leaves[0])
}
