package api

import (
	"pow-quotes-server/repositories"
)

type API struct {
	// tokensRepo - the repository for the tokens
	TokensRepo *repositories.TokensRepo

	// quotesRepo - the repository for the quotes
	QuotesRepo *repositories.QuotesRepo
}
