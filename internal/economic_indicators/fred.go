package economic_indicators

import "os"

const FRED_OBSERVATIONS_BASE_URL = "https://api.stlouisfed.org/fred/series/observations"

func FRED_API_KEY() string {
	return os.Getenv("FRED_API_KEY")
}

func getURLQuery(seriesID string) string {
	query := NewFSQuery(FRED_OBSERVATIONS_BASE_URL)
	query.And("api_key", FRED_API_KEY())
	query.And("series_id", seriesID)
	query.And("frequency", "d") // daily
	query.And("observation_start", "1950-01-01")
	return query.Build()
}