package song

import (
	sdkerrors "github.com/linden-honey/linden-honey-sdk-go/errors"
)

func (dto GetQuoteResponse) Validate() error {
	if err := dto.Quote.Validate(); err != nil {
		return sdkerrors.NewInvalidValueError("Quote", err)
	}

	return nil
}
