package vo

import (
	"fmt"
	"github.com/MikhailGulkin/blockchain-messenger/blockChain/internal/domain/services"
	"github.com/MikhailGulkin/blockchain-messenger/blockChain/internal/domain/types"
)

type MetaData struct {
	Type        types.MessageType     `json:"type"`
	Priority    types.MessagePriority `json:"priority"`
	Attachments []Attachment          `json:"attachments"`
	Location    LocationMetadata      `json:"location"`
}

func (md *MetaData) IsHighPriority() bool {
	return md.Priority == types.HighPriority
}
func (md *MetaData) AddAttachment(attachment Attachment) {
	md.Attachments = append(md.Attachments, attachment)
}
func (md *MetaData) IsLocationBased() bool {
	return md.Location != LocationMetadata{}
}

type LocationMetadata struct {
	Latitude  float64
	Longitude float64
}

func NewLocationMetadata(latitude float64, longitude float64) LocationMetadata {
	return LocationMetadata{Latitude: latitude, Longitude: longitude}
}
func (lm *LocationMetadata) DistanceTo(other LocationMetadata) string {
	distance := services.CalculateDistance(lm.Latitude, lm.Longitude, other.Latitude, other.Longitude)
	if distance == 0 {
		return "0"
	}
	return fmt.Sprintf("%.2f", distance)
}
func (lm *LocationMetadata) IsValid() bool {
	return lm.Latitude >= -90.0 && lm.Latitude <= 90.0 && lm.Longitude >= -180.0 && lm.Longitude <= 180.0
}
