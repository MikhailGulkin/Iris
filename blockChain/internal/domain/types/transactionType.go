package types

type TransactionType int

const (
	UndefinedTransaction TransactionType = iota
	MessageTransaction
)
