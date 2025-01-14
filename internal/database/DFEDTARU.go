package database

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kreimben/FinScope-engine/internal/config"
	"github.com/kreimben/FinScope-engine/internal/models"
	"github.com/kreimben/FinScope-engine/pkg/logging"
)

func SaveDFEDTARU(dfedtaru models.DFEDTARU, cfg *config.Config) error {
	query := NewSupabaseURLQuery(cfg, "economic_indicators")
	requestURL := query.Build()
	logging.Logger.WithField("url", requestURL).Debug("SAVE URL")

	economicIndicators := []models.EconomicIndicator{}

	for _, observation := range dfedtaru.Observations {
		economicIndicators = append(economicIndicators, models.EconomicIndicator{
			Name:        "DFEDTARU",
			Country:     "US",
			ReleaseDate: observation.Date.Time,
			ActualValue: observation.Value,
			Unit:        "Percent",
		})
	}

	dfedtaruJsonData, err := json.Marshal(economicIndicators)
	if err != nil {
		return err
	}

	resp, err := POST(requestURL, cfg, dfedtaruJsonData)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// check status code
	if resp.StatusCode != http.StatusCreated {
		logging.Logger.WithField("status", resp.Status).Error("Failed to save DFEDTARU")
		return errors.New("failed to save DFEDTARU")
	}

	logging.Logger.WithField("status", resp.Status).Debug("SAVE STATUS")

	return nil
}
