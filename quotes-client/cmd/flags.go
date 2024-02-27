package cmd

import (
	"flag"
	"fmt"
	"os"
)

// ClientConfig stores the configuration for the client.
type ClientConfig struct {
	BaseURL string
}

// ParseFlags parses the command line flags and returns the configuration for the client.
func ParseFlags() (*ClientConfig, error) {
	baseURL := flag.String("baseUrl", "http://localhost:8080", "Base URL of the quotes server.")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	// check if there are any flags provideds
	if len(os.Args) < 2 {
		flag.Usage()
		return nil, fmt.Errorf("no flags provided")
	}

	return &ClientConfig{
		BaseURL: *baseURL,
	}, nil
}
