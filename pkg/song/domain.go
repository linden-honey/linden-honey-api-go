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

// Song represents a domain object
type Song struct {
	ID     string  `json:"id,omitempty"`
	Title  string  `json:"title"`
	Author string  `json:"author,omitempty"`
	Album  string  `json:"album,omitempty"`
	Verses []Verse `json:"verses"`
}

// GetQuotes returns all quotes from the song
func (s Song) GetQuotes() []Quote {
	qq := make([]Quote, 0)
	for _, v := range s.Verses {
		qq = append(qq, v.Quotes...)
	}

	return qq
}

// GetRandomQuote returns a random quote from the song or an error if there are no quotes
func (s Song) GetRandomQuote() (*Quote, error) {
	r := rand.New(
		rand.NewSource(time.Now().Unix()),
	)

	qq := s.GetQuotes()
	if len(qq) == 0 {
		return nil, errors.New("no quotes")
	}
	ri := r.Intn(len(qq))

	return &qq[ri], nil
}

// GetRandomVerse returns a random verse from the song or an error if there are no verses
func (s Song) GetRandomVerse() (*Verse, error) {
	r := rand.New(
		rand.NewSource(time.Now().Unix()),
	)

	if len(s.Verses) == 0 {
		return nil, errors.New("no verses")
	}
	ri := r.Intn(len(s.Verses))

	return &s.Verses[ri], nil
}

// Preview represents a domain object
type Preview struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
