package repositories

import (
	"encoding/json"
	"errors"
	"io"
	"math/rand"

	"pow-quotes-server/models"

	"log"
	"os"

	"go.etcd.io/bbolt"
)

// QuotesRepo - the struct for the quotes repository
type QuotesRepo struct {
	db *bbolt.DB
}

// NewQuotesRepo - the constructor for the quotes repository
func NewQuotesRepo(db *bbolt.DB) *QuotesRepo {
	return &QuotesRepo{db: db}
}

// SaveQuote - saves a quote to the database
func (repo *QuotesRepo) SaveQuote(quote string, author string) error {
	err := repo.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("quotes"))
		if err != nil {
			return err
		}

		quotes := bucket.Get([]byte(quote))
		if quotes != nil {
			return nil
		}

		q := models.Quotes{Quote: quote, Author: author}
		quotes, err = json.Marshal(q)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(quote), quotes)
	})

	return err
}

// GetRandomQuote - gets a random quote from the database
func (repo *QuotesRepo) GetRandomQuote() (*models.Quotes, error) {
	var q models.Quotes

	err := repo.db.View(func(tx *bbolt.Tx) error {
		// Retrieve the quotes bucket
		bucket := tx.Bucket([]byte("quotes"))
		if bucket == nil {
			return errors.New("quotes bucket does not exist")
		}

		// Determine the total number of quotes
		stats := bucket.Stats()
		numberOfQuotes := stats.KeyN
		if numberOfQuotes == 0 {
			return errors.New("no quotes available")
		}

		// Select a random index
		randomIndex := rand.Intn(numberOfQuotes)

		c := bucket.Cursor()
		var k, v []byte
		count := 0
		// Iterate through the bucket until the random index is reached
		for k, v = c.First(); k != nil; k, v = c.Next() {
			if count == randomIndex {
				break
			}
			count++
		}

		// If k is nil, the end of the bucket was reached without finding a quote
		if k == nil {
			return errors.New("failed to find a random quote")
		}

		// Deserialize the found quote
		if err := json.Unmarshal(v, &q); err != nil {
			return err
		}

		return nil
	})

	return &q, err
}

// Function to load quotes from a JSON file into the database
func (repo *QuotesRepo) LoadQuotesIntoDB(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open quotes JSON file: %v", err)
		return err
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var quotes []models.Quotes
	if err := json.Unmarshal(byteValue, &quotes); err != nil {
		log.Fatalf("Failed to unmarshal quotes: %v", err)
		return err
	}

	// Open/create the database
	if err != nil {
		log.Fatalf("Failed to open/create the database: %v", err)
		return err
	}

	// Load the quotes into the database
	err = repo.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("quotes"))
		if err != nil {
			return err
		}

		for _, quote := range quotes {
			encoded, err := json.Marshal(quote)
			if err != nil {
				return err
			}
			// The quote itself is used as the key
			err = b.Put([]byte(quote.Quote), encoded)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Failed to load quotes into the database: %v", err)
		return err
	}
	return nil
}
