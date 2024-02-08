package song

import (
	"testing"
)

func TestGetQuoteResponse_Validate(t *testing.T) {
	type fields struct {
		Quote Quote
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Quote: "some quote",
			},
		},
		{
			name: "err  invalid quote",
			fields: fields{
				Quote: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dto := GetQuoteResponse{
				Quote: tt.fields.Quote,
			}
			if err := dto.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("GetQuoteResponse.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
