package services

import (
	"time"

	"meteo/api"
	"meteo/models"
)

func RecupererMeteo(commune models.Commune) ([]models.Day, error) {
	lon := commune.Centre.Coordinates[0]
	lat := commune.Centre.Coordinates[1]

	brut, err := api.RecupererPrevisions(lat, lon)
	if err != nil {
		return nil, err
	}

	var jours []models.Day

	for i := 0; i < len(brut.Daily.Time); i++ {
		date, _ := time.Parse("2006-01-02", brut.Daily.Time[i])

		var jour models.Day
		jour.Date = date
		jour.Min = brut.Daily.Min[i]
		jour.Max = brut.Daily.Max[i]
		jour.Code = brut.Daily.Code[i]

		jours = append(jours, jour)
	}

	return jours, nil
}
