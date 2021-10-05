package song

import (
	"errors"
	"math/rand"
	"time"
)

// Quote represents a domain object
type Quote struct {
	Phrase string `json:"phrase"`
}

// Verse represents a domain object
type Verse struct {
	Quotes []Quote `json:"quotes"`
}

// Metadata represents a domain object
type Metadata struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author,omitempty"`
	Album  string `json:"album,omitempty"`
}

// Song represents a domain object
type Song struct {
	Metadata
	Verses []Verse `json:"verses"`
}

// GetQuotes returns all quotes from the song
func (s Song) GetQuotes() []Quote {
	quotes := make([]Quote, 0)
	for _, v := range s.Verses {
		quotes = append(quotes, v.Quotes...)
	}

	return quotes
}

// GetRandomQuote returns a random quote from the song or an error if there are no quotes
func (s Song) GetRandomQuote() (*Quote, error) {
	r := rand.New(
		rand.NewSource(time.Now().Unix()),
	)

	quotes := s.GetQuotes()
	quotesCount := len(quotes)
	if quotesCount == 0 {
		return nil, errors.New("no quotes")
	}

	return &quotes[r.Intn(quotesCount)], nil
}

// GetRandomVerse returns a random verse from the song or an error if there are no verses
func (s Song) GetRandomVerse() (*Verse, error) {
	r := rand.New(
		rand.NewSource(time.Now().Unix()),
	)

	versesCount := len(s.Verses)
	if versesCount == 0 {
		return nil, errors.New("no verses")
	}

	return &s.Verses[r.Intn(versesCount)], nil
}
