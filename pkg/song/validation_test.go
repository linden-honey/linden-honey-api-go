package song

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuote_Validate(t *testing.T) {
	type fields struct {
		Phrase string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Phrase: "Some phrase",
			},
		},
		{
			name: "err  empty phrase",
			fields: fields{
				Phrase: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			q := Quote{
				Phrase: tt.fields.Phrase,
			}
			err := q.Validate()

			if tt.wantErr {
				rq.Error(err)
			} else {
				rq.NoError(err)
			}
		})
	}
}

func TestVerse_Validate(t *testing.T) {
	type fields struct {
		Quotes []Quote
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Quotes: []Quote{
					{
						Phrase: "Some phrase",
					},
				},
			},
		},
		{
			name: "err  empty quotes",
			fields: fields{
				Quotes: make([]Quote, 0),
			},
			wantErr: true,
		},
		{
			name: "err  invalid quote",
			fields: fields{
				Quotes: []Quote{
					{
						Phrase: "Some phrase",
					},
					{
						Phrase: "",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			v := Verse{
				Quotes: tt.fields.Quotes,
			}
			err := v.Validate()

			if tt.wantErr {
				rq.Error(err)
			} else {
				rq.NoError(err)
			}
		})
	}
}

func TestSong_Validate(t *testing.T) {
	type fields struct {
		Meta   Meta
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
				Meta: Meta{
					ID:     "1",
					Title:  "Some title",
					Author: "Some author",
					Album:  "Some album",
				},
				Verses: []Verse{
					{
						Quotes: []Quote{
							{
								Phrase: "Some phrase",
							},
						},
					},
				},
			},
		},
		{
			name: "err  invalid meta  empty id",
			fields: fields{
				Meta: Meta{
					ID:     "",
					Title:  "Some title",
					Author: "Some author",
					Album:  "Some album",
				},
				Verses: []Verse{
					{
						Quotes: []Quote{
							{
								Phrase: "Some phrase",
							},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "err  invalid meta  empty title",
			fields: fields{
				Meta: Meta{
					ID:     "1",
					Title:  "",
					Author: "Some author",
					Album:  "Some album",
				},
				Verses: []Verse{
					{
						Quotes: []Quote{
							{
								Phrase: "Some phrase",
							},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "err  empty verses",
			fields: fields{
				Meta: Meta{
					ID:     "1",
					Title:  "Some title",
					Author: "Some author",
					Album:  "Some album",
				},
				Verses: make([]Verse, 0),
			},
			wantErr: true,
		},
		{
			name: "err  invalid verse",
			fields: fields{
				Meta: Meta{
					ID:     "1",
					Title:  "Some title",
					Author: "Some author",
					Album:  "Some album",
				},
				Verses: []Verse{
					{
						Quotes: []Quote{
							{
								Phrase: "Some phrase",
							},
						},
					},
					{
						Quotes: make([]Quote, 0),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "err  invalid quote",
			fields: fields{
				Meta: Meta{
					ID:     "1",
					Title:  "Some title",
					Author: "Some author",
					Album:  "Some album",
				},
				Verses: []Verse{
					{
						Quotes: []Quote{
							{
								Phrase: "Some phrase",
							},
						},
					},
					{
						Quotes: []Quote{
							{
								Phrase: "",
							},
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			s := Song{
				Meta:   tt.fields.Meta,
				Verses: tt.fields.Verses,
			}
			err := s.Validate()

			if tt.wantErr {
				rq.Error(err)
			} else {
				rq.NoError(err)
			}
		})
	}
}

func TestPreview_Validate(t *testing.T) {
	type fields struct {
		ID    string
		Title string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				ID:    "123",
				Title: "Some title",
			},
		},
		{
			name: "err  empty id",
			fields: fields{
				ID:    "",
				Title: "Some title",
			},
			wantErr: true,
		},
		{
			name: "err  empty title",
			fields: fields{
				ID:    "123",
				Title: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			p := Meta{
				ID:    tt.fields.ID,
				Title: tt.fields.Title,
			}
			err := p.Validate()

			if tt.wantErr {
				rq.Error(err)
			} else {
				rq.NoError(err)
			}
		})
	}
}
