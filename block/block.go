package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/x-eight/proof_of_work-go/transaction"
)

type Block struct {
	Data         []transaction.Transaction
	time         int64
	PreviousHash string
	Hash         string
	nonce        int32
}

func NewBlock(data []transaction.Transaction, previousHash string) *Block {
	newBlock := &Block{
		Data:         data,
		time:         time.Now().Unix(),
		PreviousHash: previousHash,
		Hash:         "",
		nonce:        0,
	}
	newBlock.Hash = newBlock.CalculateHash()

	return newBlock
}

func InitString(size int) string {
	value := ""
	for i := 0; i < size; i++ {
		value = value + "0"
	}
	return value
}

func (b *Block) GetHash() string {
	return b.Hash
}

func (b *Block) CalculateHash() string {
	out, _ := json.Marshal(b.Data)

	data := string(out) + fmt.Sprint(b.time) + b.PreviousHash + fmt.Sprint(b.nonce)
	hash := sha256.Sum256([]byte(data))
	hexstring := fmt.Sprintf("%x", hash)

	return hexstring
}

func (b *Block) MineBlock(difficulty int32) int32 {

	hash := b.CalculateHash()
	for hash[0:difficulty] != InitString(int(difficulty)) {
		b.nonce = b.nonce + 1
		hash = b.CalculateHash()
	}
	b.Hash = hash
	return b.nonce
}
