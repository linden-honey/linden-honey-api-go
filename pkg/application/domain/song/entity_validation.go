package song

import (
	"errors"
	"fmt"
	"strings"

	sdkerrors "github.com/linden-honey/linden-honey-sdk-go/errors"
)

// Validate a Song and returns an error if validation is failed
func (e Entity) Validate() error {
	errs := make([]error, 0)

	if err := e.Metadata.Validate(); err != nil {
		errs = append(errs, sdkerrors.NewInvalidValueError("Metadata", err))
	}

	if err := e.Lyrics.Validate(); err != nil {
		errs = append(errs, sdkerrors.NewInvalidValueError("Lyrics", err))
	}

	return errors.Join(errs...)
}

// Validate a Metadata and returns an error if validation is failed
func (m Metadata) Validate() error {
	errs := make([]error, 0)

	if strings.TrimSpace(m.ID) == "" {
		return sdkerrors.NewRequiredValueError("ID")
	}

	if strings.TrimSpace(m.Title) == "" {
		return sdkerrors.NewRequiredValueError("Title")
	}

	if err := m.Tags.Validate(); err != nil {
		return sdkerrors.NewInvalidValueError("Tags", err)
	}

	return errors.Join(errs...)
}

// Validate Tags and returns an error if validation is failed
func (ts Tags) Validate() error {
	errs := make([]error, 0)

	unique := make(map[Tag]struct{}, len(ts))
	for _, t := range ts {
		if _, ok := unique[t]; !ok {
			unique[t] = struct{}{}
		} else {
			errs = append(
				errs,
				sdkerrors.NewInvalidValueError(fmt.Sprintf("%#v", t), errors.New("duplicate tag")),
			)

			continue
		}

		if err := t.Validate(); err != nil {
			errs = append(errs, sdkerrors.NewInvalidValueError(fmt.Sprintf("%#v", t), err))
		}
	}

	return errors.Join(errs...)
}

// Validate a Tag and returns an error if validation is failed
func (t Tag) Validate() error {
	errs := make([]error, 0)

	if strings.TrimSpace(t.Name) == "" {
		errs = append(errs, sdkerrors.NewRequiredValueError("Name"))
	}

	if strings.TrimSpace(t.Value) == "" {
		errs = append(errs, sdkerrors.NewRequiredValueError("Value"))
	}

	return errors.Join(errs...)
}

// Validate Lyrics and returns an error if validation is failed
func (l Lyrics) Validate() error {
	errs := make([]error, 0)

	if len(l) == 0 {
		errs = append(errs, sdkerrors.ErrEmptyValue)
	}

	for i, v := range l {
		if err := v.Validate(); err != nil {
			errs = append(errs, sdkerrors.NewInvalidValueError(fmt.Sprintf("Lyrics[%d]", i), err))
		}
	}

	return errors.Join(errs...)
}

// Validate a Verse and returns an error if validation is failed
func (v Verse) Validate() error {
	errs := make([]error, 0)

	if len(v) == 0 {
		errs = append(errs, sdkerrors.NewRequiredValueError("Quotes"))
	}

	for i, q := range v {
		if err := q.Validate(); err != nil {
			errs = append(errs, sdkerrors.NewInvalidValueError(fmt.Sprintf("Quotes[%d]", i), err))
		}
	}

	return errors.Join(errs...)
}

// Validate a Quote and returns an error if validation is failed
func (q Quote) Validate() error {
	if strings.TrimSpace(string(q)) == "" {
		return sdkerrors.ErrEmptyValue
	}

	return nil
}
