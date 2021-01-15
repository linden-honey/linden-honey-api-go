package song

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSong_GetQuotes(t *testing.T) {
	type fields struct {
		Verses []Verse
	}
	tests := []struct {
		name   string
		fields fields
		want   []Quote
	}{
		{
			name: "ok",
			fields: fields{
				Verses: []Verse{
					{
						Quotes: []Quote{
							{
								Phrase: "some quote",
							},
							{
								Phrase: "another quote",
							},
						},
					},
					{
						Quotes: []Quote{
							{
								Phrase: "one more quote",
							},
						},
					},
				},
			},
			want: []Quote{
				{
					Phrase: "some quote",
				},
				{
					Phrase: "another quote",
				},
				{
					Phrase: "one more quote",
				},
			},
		},
		{
			name: "empty",
			fields: fields{
				Verses: make([]Verse, 0),
			},
			want: make([]Quote, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			s := Song{
				Verses: tt.fields.Verses,
			}
			got := s.GetQuotes()

			rq.Equal(tt.want, got)
		})
	}
}

func TestSong_GetRandomQuote(t *testing.T) {
	type fields struct {
		Verses []Verse
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Verses: []Verse{
					{
						Quotes: []Quote{
							{
								Phrase: "some quote",
							},
							{
								Phrase: "another quote",
							},
						},
					},
					{
						Quotes: []Quote{
							{
								Phrase: "one more quote",
							},
						},
					},
				},
			},
		},
		{
			name: "empty",
			fields: fields{
				Verses: make([]Verse, 0),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			s := Song{
				Verses: tt.fields.Verses,
			}
			got, err := s.GetRandomQuote()

			if tt.wantErr {
				rq.Error(err)
			} else {
				rq.NoError(err)
				rq.NotNil(got)
				rq.Contains(s.GetQuotes(), *got)
			}
		})
	}
}

func TestSong_GetRandomVerse(t *testing.T) {
	type fields struct {
		Verses []Verse
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Verses: []Verse{
					{
						Quotes: []Quote{
							{
								Phrase: "some quote",
							},
							{
								Phrase: "another quote",
							},
						},
					},
					{
						Quotes: []Quote{
							{
								Phrase: "one more quote",
							},
						},
					},
				},
			},
		},
		{
			name: "empty",
			fields: fields{
				Verses: make([]Verse, 0),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			s := Song{
				Verses: tt.fields.Verses,
			}
			got, err := s.GetRandomQuote()

			if tt.wantErr {
				rq.Error(err)
			} else {
				rq.NoError(err)
				rq.NotNil(got)
				rq.Contains(s.GetQuotes(), *got)
			}
		})
	}
}
