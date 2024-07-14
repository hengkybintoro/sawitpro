package model

type (
	// Header presents as header of response
	Header struct {
		ProcessTime float64 `json:"process_time"`
	}

	// Response is data struct will be sent to end user
	Response struct {
		Header  Header      `json:"header"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}

	ResponseWithID struct {
		ID string `json:"id"`
	}

	ResponseWithStats struct {
		TreeCount    int `json:"tree_count"`
		MaxHeight    int `json:"max_height"`
		MinHeight    int `json:"min_height"`
		MedianHeight int `json:"median_height"`
	}

	ResponseWithDistance struct {
		Distance int `json:"distance"`
	}

	Coordinate struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	ResponseWithDistanceAndCoordinate struct {
		ResponseWithDistance
		Rest Coordinate `json:"rest"`
	}
)
