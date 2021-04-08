package song

import (
	"fmt"

	sdkerrors "github.com/linden-honey/linden-honey-sdk-go/errors"
)

// Validate validates a Quote and returns an error if validation is failed
func (q Quote) Validate() error {
	if q.Phrase == "" {
		return sdkerrors.NewRequiredValueError("Phrase")
	}

	return nil
}

// Validate validates a Verse and returns an error if validation is failed
func (v Verse) Validate() error {
	if len(v.Quotes) == 0 {
		return sdkerrors.NewRequiredValueError("Quotes")
	}

	for i, q := range v.Quotes {
		if err := q.Validate(); err != nil {
			return sdkerrors.NewInvalidValueError(
				fmt.Sprintf("Quotes[%d]", i),
				err,
			)
		}
	}

	return nil
}

// Validate validates a Preview and returns an error if validation is failed
func (p Preview) Validate() error {
	if p.ID == "" {
		return sdkerrors.NewRequiredValueError("ID")
	}

	if p.Title == "" {
		return sdkerrors.NewRequiredValueError("Title")
	}

	return nil
}

// Validate validates a Song and returns an error if validation is failed
func (s Song) Validate() error {
	if err := s.Preview.Validate(); err != nil {
		return err
	}

	if len(s.Verses) == 0 {
		return sdkerrors.NewRequiredValueError("Verses")
	}

	for i, v := range s.Verses {
		if err := v.Validate(); err != nil {
			return sdkerrors.NewInvalidValueError(
				fmt.Sprintf("Verses[%d]", i),
				err,
			)
		}
	}

	return nil
}
