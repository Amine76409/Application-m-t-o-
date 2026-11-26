package models

type Centre struct {
	Coordinates []float64 `json:"coordinates"`
}

type Commune struct {
	Code         string   `json:"code"`
	Nom          string   `json:"nom"`
	CodesPostaux []string `json:"codesPostaux"`
	Population   int      `json:"population"`
	Centre       struct {
		Coordinates []float64 `json:"coordinates"`
	} `json:"centre"`
}
