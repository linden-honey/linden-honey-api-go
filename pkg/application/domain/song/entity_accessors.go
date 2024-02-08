package song

import (
	"errors"
	"math/rand"
)

// GetRandomVerse returns a random verse from lyrics or an error if there are no verses.
func (l Lyrics) GetRandomVerse() (Verse, error) {
	versesCount := len(l)
	if versesCount == 0 {
		return nil, errors.New("no verses")
	}

	return l[rand.Intn(versesCount)], nil
}

// GetRandomQuote returns a random quote from lyrics or an error if there are no quotes.
func (l Lyrics) GetRandomQuote() (Quote, error) {
	quotes := l.GetQuotes()
	quotesCount := len(quotes)
	if quotesCount == 0 {
		return "", errors.New("no quotes")
	}

	return quotes[rand.Intn(quotesCount)], nil
}

// GetQuotes returns all quotes from the song.
func (l Lyrics) GetQuotes() []Quote {
	quotes := make([]Quote, 0)
	for _, v := range l {
		quotes = append(quotes, v...)
	}

	return quotes
}
