package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Chain []*Block

var BlockChain Chain

// Block represents an individual block on the blockchain
type Block struct {
	Timestamp    int64         `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	Proof        int           `json:"proof"`
	PreviousHash string        `json:"previous_hash"`
}

// Transaction represents a transfer of funds from one user to another
type Transaction struct {
	Sender    string
	Recipient string
	Amount    float32
}

// NewBlock returns a fresh Block
func NewBlock(proof int, previousHash string) *Block {

	if previousHash == "" {
		previousHash = Hash(BlockChain.GetLastBlock())
	}

	block := &Block{
		Timestamp:    time.Now().UnixNano(),
		Proof:        proof,
		PreviousHash: previousHash}

	BlockChain = append(BlockChain, block)

	return block
}

// NewTransaction adds a new transaction to the block
func (b *Block) NewTransaction(sender string, recipient string, amount float32) {

	t := Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount}

	b.Transactions = append(b.Transactions, t)

}

// Mine performs a proof of work and writes the pending transactions to the blockchain when a proof is achieved
func (b *Block) Mine() *Block {
	lastBlock := BlockChain.GetLastBlock()
	lastProof := lastBlock.Proof

	newProof := proofOfWork(lastProof)

	// Reward ourselves
	b.NewTransaction("0", "ourselves", 1)

	nb := NewBlock(newProof, "")

	return nb
}

func Hash(block *Block) string {
	// Convert block into JSON
	jb, err := json.Marshal(block)
	if err != nil {
		panic(err)
	}

	// Compute and return SHA-256 of the Block as hex
	h := sha256.New()
	h.Write([]byte(jb))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (c Chain) GetLastBlock() *Block {
	if len(c) == 0 {
		return nil
	}
	return c[len(c)-1]
}

func isValidProof(lastProof int, cur int) bool {
	guess := fmt.Sprintf("%s%s", lastProof, cur)

	h := sha256.New()
	h.Write([]byte(guess))
	gh := fmt.Sprintf("%x", h.Sum(nil))
	last4 := gh[len(gh)-4:]

	return last4 == "0000"
}

func proofOfWork(lastProof int) int {
	cur := 0
	for {
		if isValidProof(lastProof, cur) {
			return cur
		}
		cur++
	}
}
