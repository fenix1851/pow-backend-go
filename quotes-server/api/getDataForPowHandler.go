package api

import (
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *API) GetDataForPowHandler(c *gin.Context) {
	// create a random number - it will be used as the data for the PoW
	randomNumber := int64(rand.Intn(999999))

	// create a random string - it will be used as the PoW token
	randomString := randomString(16)

	// take the leading zeros count from the cache
	leadingZerosCount, error := api.TokensRepo.GetCurrentLeadingZerosCount()
	if error != nil {
		c.JSON(500, gin.H{
			"error": error.Error(),
		})
		return
	}

	error = api.TokensRepo.SaveToken(randomString, randomNumber, leadingZerosCount)
	if error != nil {
		c.JSON(500, gin.H{
			"error": error.Error(),
		})
		return
	}

	// return the data and the token
	c.JSON(200, gin.H{
		"data":              randomNumber,
		"token":             randomString,
		"leadingZerosCount": leadingZerosCount,
	})

}

func randomString(length int) string {
	var result string
	for len(result) < length {
		// Add a random string to the result
		result += strconv.FormatInt(rand.Int63(), 36)
	}

	// Cut the result to the required length
	return result[:length]
}
