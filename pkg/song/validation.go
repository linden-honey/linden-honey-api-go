package song

import (
	"errors"
	"fmt"
	"strings"

	sdkerrors "github.com/linden-honey/linden-honey-sdk-go/errors"
)

// Validate a Tag and returns an error if validation is failed
func (t Tag) Validate() error {
	if strings.TrimSpace(t.Name) == "" {
		return sdkerrors.NewRequiredValueError("Name")
	}

	if strings.TrimSpace(t.Value) == "" {
		return sdkerrors.NewRequiredValueError("Value")
	}

	return nil
}

// Validate Tags and returns an error if validation is failed
func (ts Tags) Validate() error {
	unique := make(map[Tag]struct{}, len(ts))
	for i, t := range ts {
		if _, ok := unique[t]; !ok {
			unique[t] = struct{}{}
		} else {
			return sdkerrors.NewInvalidValueError(fmt.Sprintf("Tags[%d]", i), errors.New("duplicate tag"))
		}

		if err := t.Validate(); err != nil {
			return sdkerrors.NewInvalidValueError(fmt.Sprintf("Tags[%d]", i), err)
		}
	}

	return nil
}

// Validate a Metadata and returns an error if validation is failed
func (m Metadata) Validate() error {
	if strings.TrimSpace(m.ID) == "" {
		return sdkerrors.NewRequiredValueError("ID")
	}

	if strings.TrimSpace(m.Title) == "" {
		return sdkerrors.NewRequiredValueError("Title")
	}

	if err := m.Tags.Validate(); err != nil {
		return sdkerrors.NewInvalidValueError("Tags", err)
	}

	return nil
}

// Validate a Quote and returns an error if validation is failed
func (q Quote) Validate() error {
	if strings.TrimSpace(q.Phrase) == "" {
		return sdkerrors.NewRequiredValueError("Phrase")
	}

	return nil
}

// Validate a Verse and returns an error if validation is failed
func (v Verse) Validate() error {
	if len(v.Quotes) == 0 {
		return sdkerrors.NewRequiredValueError("Quotes")
	}

	for i, q := range v.Quotes {
		if err := q.Validate(); err != nil {
			return sdkerrors.NewInvalidValueError(fmt.Sprintf("Quotes[%d]", i), err)
		}
	}

	return nil
}

// Validate Lyrics and returns an error if validation is failed
func (l Lyrics) Validate() error {
	if len(l) == 0 {
		return sdkerrors.ErrEmptyValue
	}

	for i, v := range l {
		if err := v.Validate(); err != nil {
			return sdkerrors.NewInvalidValueError(fmt.Sprintf("Lyrics[%d]", i), err)
		}
	}

	return nil
}

// Validate a Song and returns an error if validation is failed
func (s Song) Validate() error {
	if err := s.Metadata.Validate(); err != nil {
		return err
	}

	if err := s.Lyrics.Validate(); err != nil {
		return sdkerrors.NewInvalidValueError("Lyrics", err)
	}

	return nil
}
