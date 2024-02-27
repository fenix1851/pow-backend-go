package cmd

import (
	"flag"
	"fmt"
	"os"
	"pow-quotes-server/config"
)

// AppConfig contains the application configuration.

func ParseFlags() (*config.Config, error) {
	quotesJsonPath := flag.String("quotesJsonPath", "", "Path to the JSON file containing the quotes.")
	timeFrame := flag.Int64("timeFrame", 20, "Time frame in seconds for getting the number of requests to recalculate the leading zeros count.")
	requestsThreshold := flag.Int64("requestsThreshold", 1, "The number of requests in a given time frame will be divided by this number to get the number of leading zeros for the PoW.")
	leadingZerosUpdateInterval := flag.Int64("leadingZerosUpdateInterval", 1, "The interval in seconds for updating the leading zeros count.")
	port := flag.String("port", "8080", "Port to run the server on.")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	// Check if no flags were provided and display usage if true
	if len(os.Args) < 2 {
		flag.Usage()
		return nil, fmt.Errorf("no flags provided")
	}

	return &config.Config{
		QuotesJsonPath:             *quotesJsonPath,
		TimeFrame:                  *timeFrame,
		RequestsThreshold:          *requestsThreshold,
		LeadingZerosUpdateInterval: *leadingZerosUpdateInterval,
		Port:                       *port,
		Cache: config.Cache{
			LeadingZerosCountCache: &config.LeadingZerosCountCache{},
		},
	}, nil

}
