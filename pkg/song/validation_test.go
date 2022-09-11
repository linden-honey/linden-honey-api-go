package song

import (
	"testing"
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
				Phrase: "some phrase",
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
			q := Quote{
				Phrase: tt.fields.Phrase,
			}
			if err := q.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Quote.Validate() error = %v, wantErr %v", err, tt.wantErr)
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
						Phrase: "some phrase",
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
						Phrase: "some phrase",
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
			v := Verse{
				Quotes: tt.fields.Quotes,
			}
			if err := v.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Verse.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMetadata_Validate(t *testing.T) {
	type fields struct {
		ID     string
		Title  string
		Author string
		Album  string
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
				Title: "some title",
			},
		},
		{
			name: "err  empty id",
			fields: fields{
				ID:    "",
				Title: "some title",
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
			p := Metadata{
				ID:     tt.fields.ID,
				Title:  tt.fields.Title,
				Author: tt.fields.Author,
				Album:  tt.fields.Album,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Metadata.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSong_Validate(t *testing.T) {
	type fields struct {
		Metadata Metadata
		Verses   []Verse
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Metadata: Metadata{
					ID:     "1",
					Title:  "some title",
					Author: "some author",
					Album:  "some album",
				},
				Verses: []Verse{
					{
						Quotes: []Quote{
							{
								Phrase: "some phrase",
							},
						},
					},
				},
			},
		},
		{
			name: "err  invalid meta  empty id",
			fields: fields{
				Metadata: Metadata{
					ID:     "",
					Title:  "some title",
					Author: "some author",
					Album:  "some album",
				},
				Verses: []Verse{
					{
						Quotes: []Quote{
							{
								Phrase: "some phrase",
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
				Metadata: Metadata{
					ID:     "1",
					Title:  "",
					Author: "some author",
					Album:  "some album",
				},
				Verses: []Verse{
					{
						Quotes: []Quote{
							{
								Phrase: "some phrase",
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
				Metadata: Metadata{
					ID:     "1",
					Title:  "some title",
					Author: "some author",
					Album:  "some album",
				},
				Verses: make([]Verse, 0),
			},
			wantErr: true,
		},
		{
			name: "err  invalid verse",
			fields: fields{
				Metadata: Metadata{
					ID:     "1",
					Title:  "some title",
					Author: "some author",
					Album:  "some album",
				},
				Verses: []Verse{
					{
						Quotes: []Quote{
							{
								Phrase: "some phrase",
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
				Metadata: Metadata{
					ID:     "1",
					Title:  "some title",
					Author: "some author",
					Album:  "some album",
				},
				Verses: []Verse{
					{
						Quotes: []Quote{
							{
								Phrase: "some phrase",
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
			s := Song{
				Metadata: tt.fields.Metadata,
				Verses:   tt.fields.Verses,
			}
			if err := s.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Song.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
