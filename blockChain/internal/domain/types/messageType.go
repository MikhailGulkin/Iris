package types

type MessageType int

const (
	UndefinedMessage MessageType = iota
	TextMessage
	ImageMessage
	VideoMessage
	AudioMessage
)

type MessagePriority int

const (
	UndefinedPriority MessagePriority = iota
	NormalPriority
	HighPriority
	LowPriority
)
