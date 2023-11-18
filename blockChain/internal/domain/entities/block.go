package entities

import (
	"crypto/sha256"
	"encoding/json"
)

type Block struct {
	Id           string         `json:"id"`
	PrevHash     [32]byte       `json:"prevHash"`
	Transactions []*Transaction `json:"transactions"`
	TimeStamp    string         `json:"timeStamp"`
	Nonce        int64          `json:"nonce"`
}

func NewBlock(once int64, prevHash [32]byte, transactions []*Transaction) *Block {
	panic("implement me")
}
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256(m)
}
