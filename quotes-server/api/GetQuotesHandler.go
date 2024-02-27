package api

import "github.com/gin-gonic/gin"

func (api *API) GetQuotesHandler(c *gin.Context) {
	quotes, err := api.QuotesRepo.GetRandomQuote()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, quotes)
}
