package vo

type Attachment struct {
	URL string `json:"url"`
}

func NewAttachment(url string) Attachment {
	return Attachment{URL: url}
}
