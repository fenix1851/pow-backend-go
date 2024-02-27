package main

import (
	"fmt"
	"os"
	"pow-quotes-server/api"
	"pow-quotes-server/cmd"
	"pow-quotes-server/config"
	"pow-quotes-server/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	flagsConfig, err := cmd.ParseFlags()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db := config.CreateDB()

	api := api.API{
		TokensRepo: repositories.NewTokensRepo(db),
		QuotesRepo: repositories.NewQuotesRepo(db),
	}
	// Validate the parsed config
	err = config.ValidateConfig(flagsConfig)
	if err != nil {
		fmt.Printf("Failed to create app config: %v\n", err)
		return
	}

	// set the global variable AppConfig to the parsed config
	config.AppConfig = *flagsConfig

	// Load the quotes into the database if the path to the JSON file is provided
	if config.AppConfig.QuotesJsonPath != "" {
		err := api.QuotesRepo.LoadQuotesIntoDB(config.AppConfig.QuotesJsonPath)
		if err != nil {
			fmt.Printf("Failed to load quotes into the database: %v\n", err)
			return
		}
	}

	// Start the server
	r := gin.Default()

	// Routes
	r.GET("/api/v1/pow/", api.GetDataForPowHandler)
	r.GET("/api/v1/quotes/", api.PoWValidationMiddleware, api.GetQuotesHandler)

	// Run the server
	fmt.Printf("Server is running at :%s\n", config.AppConfig.Port)
	r.Run(":" + config.AppConfig.Port)
}
