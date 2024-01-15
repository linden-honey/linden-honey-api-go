package song

import (
	"errors"
	"math/rand"
)

// Entity is a main domain object.
type Entity struct {
	Metadata
	Lyrics `json:"lyrics"`
}

// Metadata is a domain helper object.
type Metadata struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Tags  Tags   `json:"tags,omitempty"`
}

// Tags is a domain object.
type Tags []Tag

// Tag is a domain object.
type Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Lyrics is a domain object.
type Lyrics []Verse

// GetRandomVerse returns a random verse from lyrics or an error if there are no verses.
func (l Lyrics) GetRandomVerse() (*Verse, error) {
	versesCount := len(l)
	if versesCount == 0 {
		return nil, errors.New("no verses")
	}

	return &l[rand.Intn(versesCount)], nil
}

// GetRandomQuote returns a random quote from lyrics or an error if there are no quotes.
func (l Lyrics) GetRandomQuote() (*Quote, error) {
	quotes := l.GetQuotes()
	quotesCount := len(quotes)
	if quotesCount == 0 {
		return nil, errors.New("no quotes")
	}

	return &quotes[rand.Intn(quotesCount)], nil
}

// GetQuotes returns all quotes from the song.
func (l Lyrics) GetQuotes() []Quote {
	quotes := make([]Quote, 0)
	for _, v := range l {
		quotes = append(quotes, v.Quotes...)
	}

	return quotes
}

// Verse is a domain object.
type Verse struct {
	Quotes []Quote `json:"quotes"`
}

// Quote is a domain object.
type Quote struct {
	Phrase string `json:"phrase"`
}
