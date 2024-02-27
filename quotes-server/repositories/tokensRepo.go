package repositories

import (
	"encoding/json"

	"pow-quotes-server/config"
	"pow-quotes-server/models"
	"time"

	"go.etcd.io/bbolt"
)

// TokensRepo - the struct for the tokens repository
type TokensRepo struct {
	db *bbolt.DB
}

// NewTokensRepo - the constructor for the tokens repository
func NewTokensRepo(db *bbolt.DB) *TokensRepo {
	return &TokensRepo{db: db}
}

// SaveToken - saves a token to the database
func (repo *TokensRepo) SaveToken(token string, data int64, leadingZerosCount int64) error {
	err := repo.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("tokens"))
		if err != nil {
			return err
		}

		tokens := bucket.Get([]byte(token))
		if tokens != nil {
			return nil
		}

		t := models.Tokens{Token: token, Data: data, LeadingZerosCount: leadingZerosCount, Timestamp: time.Now().Unix()}
		tokens, err = json.Marshal(t)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(token), tokens)
	})

	return err
}

// GetToken - gets a token from the database
func (repo *TokensRepo) GetToken(token string) (*models.Tokens, error) {
	var t models.Tokens

	err := repo.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("tokens"))
		if bucket == nil {
			return nil
		}

		tokens := bucket.Get([]byte(token))
		if tokens == nil {
			return nil
		}

		return json.Unmarshal(tokens, &t)
	})

	if err != nil {
		return nil, err
	}

	return &t, nil
}

// Get tokens count in period
func (repo *TokensRepo) GetTokensCountInPeriod(from int64, to int64) (int64, error) {
	var count int64

	err := repo.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("tokens"))
		if bucket == nil {
			return nil
		}

		return bucket.ForEach(func(k, v []byte) error {
			var t models.Tokens
			err := json.Unmarshal(v, &t)
			if err != nil {
				return err
			}

			if t.Timestamp >= from && t.Timestamp <= to {
				count++
			}

			return nil
		})
	})

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (repo *TokensRepo) GetCurrentLeadingZerosCount() (int64, error) {
	var leadingZerosCount int64
	var lastUpdate int64
	var err error

	// get the current time
	currentTime := time.Now().Unix()

	// get the last update time from the cache
	lastUpdate, err = config.AppConfig.Cache.GetLastUpdate()
	if err != nil {
		return 0, err
	}

	// if the last update was more than the update interval, recalculate the leading zeros count
	// set the last update time and save it to the cache
	if currentTime-lastUpdate > config.AppConfig.LeadingZerosUpdateInterval {
		// get the tokens count in the last time frame
		// last time frame is the current time minus the TimeFrame from the config
		tokensCountInLastPeriod, err := repo.GetTokensCountInPeriod(currentTime-config.AppConfig.TimeFrame, currentTime)
		if err != nil {
			return 0, err
		}
		// get the leading zeros count from the tokens count
		// the leading zeros count is the tokens count divided by the RequestsThreshold from the config
		leadingZerosCount = tokensCountInLastPeriod / config.AppConfig.RequestsThreshold

		// set the leading zeros count and the last update time to the cache
		err = config.AppConfig.Cache.SetLeadingZerosCount(leadingZerosCount, currentTime)
		if err != nil {
			return 0, err
		}

		err = config.AppConfig.Cache.SetLastUpdate(currentTime)
		if err != nil {
			return 0, err
		}
		return leadingZerosCount, nil
	}

	// get the leading zeros count from the cache
	leadingZerosCount, _, err = config.AppConfig.Cache.GetLeadingZerosCount()
	if err != nil {
		return 0, err
	}

	return leadingZerosCount, nil
}
