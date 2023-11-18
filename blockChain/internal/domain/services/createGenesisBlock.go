package services

import "github.com/MikhailGulkin/blockchain-messenger/blockChain/internal/domain/entities"

func NewGenesisBlock() *entities.Block {
	return &entities.Block{}
}
