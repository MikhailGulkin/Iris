package entities

import (
	"github.com/MikhailGulkin/blockchain-messenger/blockChain/internal/domain/types"
	"github.com/MikhailGulkin/blockchain-messenger/blockChain/internal/domain/vo"
	"time"
)

type Transaction struct {
	Id              vo.TransactionID      `json:"id"`
	FromAddress     vo.Address            `json:"fromAddress"`
	ToAddress       []vo.Address          `json:"toAddress"`
	Message         vo.Message            `json:"message"`
	Amount          vo.Amount             `json:"amount"`
	BlockID         vo.BlockID            `json:"blockId"`
	Signature       string                `json:"signature"`
	TransactionType types.TransactionType `json:"transactionType"`
	TimeStamp       time.Time             `json:"timeStamp"`
}

func NewTransaction() {
	return
}
