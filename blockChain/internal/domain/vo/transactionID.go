package vo

import "github.com/google/uuid"

type TransactionID struct {
	value uuid.UUID
}

func NewTransactionID() TransactionID {
	return TransactionID{value: uuid.New()}
}
