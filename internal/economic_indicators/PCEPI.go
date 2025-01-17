package economic_indicators

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/kreimben/FinScope-engine/internal/config"
	"github.com/kreimben/FinScope-engine/internal/database"
	"github.com/kreimben/FinScope-engine/internal/models"
	"github.com/kreimben/FinScope-engine/pkg/logging"
)

const PCEPI_NAME = "PCEPI"
const PCEPI_OBSERVATION_START_DATE = "1960-01-01"

func GatherPCEPI(cfg *config.Config) error {
	logging.Logger.Debug("Getting PCEPI")

	url := getFREDQuery(PCEPI_NAME, PCEPI_OBSERVATION_START_DATE, "m", "pc1")
	logging.Logger.WithField("url", url).Debug("FRED URL")

	response, err := http.Get(url)
	if err != nil {
		logging.Logger.WithError(err).Error("Error getting PCEPI")
		return err
	}
	defer response.Body.Close()

	var pcepi models.PCEPI
	body, err := io.ReadAll(response.Body)
	if err != nil {
		logging.Logger.WithError(err).Error("Error reading PCEPI response body")
		return err
	}
	err = json.Unmarshal(body, &pcepi)
	if err != nil {
		logging.Logger.WithError(err).Error("Error unmarshalling PCEPI")
		return err
	}

	database.DeleteAllEconomicIndicators(cfg, PCEPI_NAME)

	return database.SavePCEPI(pcepi, cfg)
}
