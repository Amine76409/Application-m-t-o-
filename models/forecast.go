package models

import "time"

type Day struct {
	Date time.Time
	Min  float64
	Max  float64
	Code int
}

type DonneesJours struct {
	Time []string  `json:"time"`
	Max  []float64 `json:"temperature_2m_max"`
	Min  []float64 `json:"temperature_2m_min"`
	Code []int     `json:"weather_code"`
}

type ReponseMeteo struct {
	Daily DonneesJours `json:"daily"`
}
