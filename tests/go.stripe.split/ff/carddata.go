package stripe

type CardData struct {
	Object string  `json:"object"`
	Count  int     `json:"count"`
	Url    string  `json:"url"`
	Data   []*Card `json:"data"`
}
