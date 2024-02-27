package config

import "errors"

type Config struct {
	// timeFrame - the time frame in seconds for getting the number of requests to recalculate the leading zeros count
	TimeFrame int64

	// requestsThreshold - the number of requests in a given time frame will be devided by this number,
	// to get the number of leading zeros for the PoW
	RequestsThreshold int64

	// LeadingZerosUpdateInterval - the interval in seconds for updating the leading zeros count
	LeadingZerosUpdateInterval int64

	//port - the port to run the server on
	Port string

	//QuotesJsonPath - the path to the JSON file containing the quotes
	QuotesJsonPath string

	Cache Cache
}

func ValidateConfig(config *Config) error {
	// validate the input
	if config.TimeFrame <= 0 {
		return errors.New("timeFrame must be greater than 0")
	}

	if config.RequestsThreshold <= 0 {
		return errors.New("requestsThreshold must be greater than 0")
	}

	if config.LeadingZerosUpdateInterval <= 0 {
		return errors.New("leadingZerosUpdateInterval must be greater than 0")
	}
	return nil
}

var AppConfig Config
