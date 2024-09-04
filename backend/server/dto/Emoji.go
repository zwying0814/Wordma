package dto

type EmojiJson struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Images []struct {
		Icon string `json:"icon"`
		Flag string `json:"flag"`
		Text string `json:"text"`
	} `json:"images"`
}
