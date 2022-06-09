package main

import (
	"fmt"

	"github.com/x-eight/proof_of_work-go/block"
	"github.com/x-eight/proof_of_work-go/blockchain"
	"github.com/x-eight/proof_of_work-go/key"
	"github.com/x-eight/proof_of_work-go/transaction"
)

func main() {
	mint := *key.Init()

	friendWallet := key.GenKeyPairT()
	holder := *key.GenKeyPairT()

	red := blockchain.NewBlockChain()

	//=================give crypto new user =======================//
	tx1 := transaction.NewTransaction(mint.GET_MINT_PUBLIC_ADDRESS(), holder.Public, 100000)
	red.AddTransaction(*tx1)
	hash := red.GetLatestBlock().Hash
	newBlock := block.NewBlock([]transaction.Transaction{*tx1}, hash)
	red.AddBlock(*newBlock)

	//=================give to preseent crypto my friend =======================//
	tx2 := transaction.NewTransaction(holder.Public, friendWallet.Public, 333, 10)
	tx2.Sign(holder)

	balance := red.GetBalance(tx2.FromAddress)
	if tx2.IsValid(balance, red.Reward, holder, mint.GET_MINT_PUBLIC_ADDRESS()) {
		red.AddTransaction(*tx2)

	}

	red.MineTransactions(holder.Public, mint)

	fmt.Println("Your balance:")
	fmt.Println(red.GetBalance(holder.Public))
	fmt.Println("friend's balance:")
	fmt.Println(red.GetBalance(friendWallet.Public))

}
