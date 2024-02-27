package api

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PoWValidationRequest struct {
	Token string `json:"token"`
	Nonce int64  `json:"nonce"`
}

func (api *API) PoWValidationMiddleware(c *gin.Context) {
	// bind the get parameters to the struct
	if c.Query("token") == "" || c.Query("nonce") == "" {
		c.JSON(400, gin.H{
			"error": "The token and the nonce are required",
		})
		return
	}
	var request PoWValidationRequest
	request.Token = c.Query("token")
	request.Nonce, _ = strconv.ParseInt(c.Query("nonce"), 10, 64)
	// get the token from the database
	token, err := api.TokensRepo.GetToken(request.Token)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	// check if the hash is correct
	if !isHashCorrect(token.Data, request.Nonce, token.LeadingZerosCount) {
		c.JSON(400, gin.H{
			"error": "The nonce is not correct",
		})
		return
	}
	c.Next()
}

func isHashCorrect(data int64, nonce int64, leadingZerosCount int64) bool {
	// create the string that will be hashed
	stringToHash := strconv.FormatInt(data, 10) + strconv.FormatInt(nonce, 10)

	// create the hash
	hash := sha256.New()
	hash.Write([]byte(stringToHash))
	hashed := hex.EncodeToString(hash.Sum(nil))

	// check if the hash is correct
	for i := int64(0); i < leadingZerosCount; i++ {
		if hashed[i] != '0' {
			return false
		}
	}
	return true
}
