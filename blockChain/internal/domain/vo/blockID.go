package vo

import "github.com/google/uuid"

type BlockID struct {
	value uuid.UUID
}

func NewBlockID() BlockID {
	return BlockID{value: uuid.New()}
}
