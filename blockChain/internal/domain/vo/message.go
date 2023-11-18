package vo

import "time"

type Message struct {
	Text      string    `json:"message"`
	TimeStamp time.Time `json:"timeStamp"`
	IsRead    bool      `json:"isRead"`
	MetaData  MetaData  `json:"metaData"`
}
