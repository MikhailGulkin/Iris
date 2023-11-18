package vo

import "github.com/MikhailGulkin/blockchain-messenger/blockChain/internal/domain/types"

type MetaData struct {
	Type        types.MessageType     `json:"type"`
	Priority    types.MessagePriority `json:"priority"`
	Attachments []Attachment          `json:"attachments"`
	Location    LocationMetadata      `json:"location"`
}
type LocationMetadata struct {
	Latitude  float64
	Longitude float64
}
