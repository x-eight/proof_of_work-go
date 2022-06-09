package transaction

import (
	"fmt"

	"github.com/x-eight/proof_of_work-go/key"
)

type Transaction struct {
	FromAddress string
	ToAddress   string
	Amount      int
	Gas         int
	Signature   string
}

func NewTransaction(from string, to string, amount int, gas_op ...int) *Transaction {
	gas := 0
	if len(gas_op) != 0 {
		gas = gas_op[0]
	}

	return &Transaction{
		FromAddress: from,
		ToAddress:   to,
		Amount:      amount,
		Gas:         gas,
		Signature:   "",
	}
}

func (t *Transaction) Sign(key key.Key) {
	if key.Public == t.FromAddress {
		// Add gas
		t.Signature, _ = key.Sign(t.FromAddress, t.ToAddress, fmt.Sprint(t.Amount), fmt.Sprint(t.Gas))
	}
}

func (t *Transaction) IsValid(balanceAmaunt, reward int, key key.Key, address_op ...string) bool {
	address := "MINT_PUBLIC_ADDRESS"
	if len(address_op) != 0 {
		address = address_op[0]
	}

	balance := (balanceAmaunt >= t.Amount+t.Gas) || (t.FromAddress == address && t.Amount == reward)
	signature := key.Verify(t.FromAddress, t.Signature, t.FromAddress, t.ToAddress, fmt.Sprint(t.Amount), fmt.Sprint(t.Gas))
	return (t.FromAddress != "" && t.ToAddress != "" && t.Amount != 0 && balance && signature)

}
