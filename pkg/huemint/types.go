package huemint

type reqPayload struct {
	NumColors   int      `json:"num_colors"`
	Temperature float64  `json:"temperature"`
	NumResults  int      `json:"num_results"`
	Adjacency   []string `json:"adjacency"`
	Palette     []string `json:"palette"`
	Mode        string   `json:"mode"`
}

type respData struct {
	Results []struct {
		Palette []string `json:"palette"`
		Score   float64  `json:"score"`
	} `json:"results"`
}
