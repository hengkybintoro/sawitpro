package model

type (
	// Request is data struct will be received from user
	RequestEstate struct {
		Length int `json:"length"`
		Width  int `json:"width"`
	}

	RequestTree struct {
		X      int `json:"x"`
		Y      int `json:"y"`
		Height int `json:"height"`
	}
)
