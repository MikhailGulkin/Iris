package vo

import "time"

type Message struct {
	Text      string    `json:"message"`
	IsRead    bool      `json:"isRead"`
	MetaData  MetaData  `json:"metaData"`
	TimeStamp time.Time `json:"timeStamp"`
}
