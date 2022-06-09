package blockchain

import (
	"github.com/x-eight/proof_of_work-go/block"
	"github.com/x-eight/proof_of_work-go/key"
	"github.com/x-eight/proof_of_work-go/transaction"
)

type BlockChain struct {
	Chain        []block.Block
	Difficulty   int32
	BlockTime    int
	Transactions []transaction.Transaction
	Reward       int
}

func NewBlockChain() *BlockChain {

	genesis := *block.NewBlock([]transaction.Transaction{}, "00000000000000000000000000000000")

	return &BlockChain{
		Chain:        []block.Block{genesis},
		Difficulty:   1,
		BlockTime:    30000,
		Transactions: []transaction.Transaction{},
		Reward:       0,
	}
}

func (b *BlockChain) GetBalance(address string) int {
	balance := 0
	for _, block := range b.Chain {
		for _, tx := range block.Data {

			if tx.FromAddress == address {
				balance = balance - tx.Amount - tx.Gas
			}
			if tx.ToAddress == address {
				balance = balance + tx.Amount
			}

		}
	}
	return balance
}

func (b *BlockChain) GetLatestBlock() block.Block {
	return b.Chain[len(b.Chain)-1]
}

func (b *BlockChain) AddBlock(newBlock block.Block) {
	newBlock.MineBlock(b.Difficulty)

	b.Chain = append(b.Chain, newBlock)
	b.Transactions = []transaction.Transaction{}
}

func (b *BlockChain) IsChainValid() bool {

	for i := 1; i < len(b.Chain); i++ {
		currentBlock := b.Chain[i]
		previousBlock := b.Chain[i-1]

		if currentBlock.Hash != currentBlock.CalculateHash() {
			return false
		}

		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}

	return true
}

func (b *BlockChain) AddTransaction(tx transaction.Transaction) {
	b.Transactions = append(b.Transactions, tx)
}

func (b *BlockChain) MineTransactions(rewardAddress string, mint key.Mint) {
	rewardTx := transaction.NewTransaction(mint.GET_MINT_PUBLIC_ADDRESS(), rewardAddress, b.Reward)
	rewardTx.Sign(mint.GET_MINT_KEY_PAIR())

	transactions := append(b.Transactions, *rewardTx)

	if len(b.Transactions) != 0 {
		lastBlock := b.GetLatestBlock()
		test := *block.NewBlock(transactions, lastBlock.Hash)
		b.AddBlock(test)

	}

	b.Transactions = []transaction.Transaction{}
}
