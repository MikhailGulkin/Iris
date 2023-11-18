package aggregate

import (
	"github.com/MikhailGulkin/blockchain-messenger/blockChain/internal/domain/entities"
	"github.com/MikhailGulkin/blockchain-messenger/blockChain/internal/domain/services"
	"github.com/MikhailGulkin/blockchain-messenger/blockChain/internal/domain/vo"
)

type BlockChain struct {
	Chain               []*entities.Block       `json:"chain"`
	PendingTransactions []*entities.Transaction `json:"pendingTransactions"`
	BlockChainAddress   vo.Address              `json:"blockChainAddress"`
}

func NewBlockChain(address vo.Address) *BlockChain {
	b := services.NewGenesisBlock()
	bc := new(BlockChain)
	bc.BlockChainAddress = address
	bc.InsertBlock(0, b.Hash())
	return bc
}

func (bc *BlockChain) InsertBlock(once int64, prevHash [32]byte) *entities.Block {
	b := entities.NewBlock(once, prevHash, bc.PendingTransactions)
	bc.PendingTransactions = []*entities.Transaction{}
	bc.Chain = append(bc.Chain, b)
	return b
}
