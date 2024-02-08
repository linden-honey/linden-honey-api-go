package song

import (
	"testing"
)

func TestSong_Validate(t *testing.T) {
	type fields struct {
		Metadata Metadata
		Lyrics   Lyrics
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
					ID:    "1",
					Title: "some title",
					Tags: Tags{
						Tag{
							Name:  "author",
							Value: "some author",
						},
						Tag{
							Name:  "album",
							Value: "some album",
						},
					},
				},
				Lyrics: Lyrics{
					{
						"some quote",
					},
				},
			},
		},
		{
			name: "err  invalid metadata",
			fields: fields{
				Metadata: Metadata{
					ID:    "",
					Title: "some title",
				},
				Lyrics: Lyrics{
					{
						"some quote",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "err  invalid lyrics",
			fields: fields{
				Metadata: Metadata{
					ID:    "1",
					Title: "some title",
					Tags: Tags{
						Tag{
							Name:  "author",
							Value: "some author",
						},
						Tag{
							Name:  "album",
							Value: "some album",
						},
					},
				},
				Lyrics: Lyrics{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Entity{
				Metadata: tt.fields.Metadata,
				Lyrics:   tt.fields.Lyrics,
			}
			if err := s.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Song.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMetadata_Validate(t *testing.T) {
	type fields struct {
		ID    string
		Title string
		Tags  Tags
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
				Tags: Tags{
					Tag{
						Name:  "author",
						Value: "some author",
					},
					Tag{
						Name:  "album",
						Value: "some album",
					},
				},
			},
		},
		{
			name: "err  empty id",
			fields: fields{
				ID:    "",
				Title: "some title",
				Tags: Tags{
					Tag{
						Name:  "author",
						Value: "some author",
					},
					Tag{
						Name:  "album",
						Value: "some album",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "err  empty title",
			fields: fields{
				ID:    "123",
				Title: "",
				Tags: Tags{
					Tag{
						Name:  "author",
						Value: "some author",
					},
					Tag{
						Name:  "album",
						Value: "some album",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "err  invalid tags",
			fields: fields{
				ID:    "123",
				Title: "some title",
				Tags: Tags{
					Tag{
						Name:  "author",
						Value: "some author",
					},
					Tag{
						Name:  "album",
						Value: "",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Metadata{
				ID:    tt.fields.ID,
				Title: tt.fields.Title,
				Tags:  tt.fields.Tags,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Metadata.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTags_Validate(t *testing.T) {
	tests := []struct {
		name    string
		ts      Tags
		wantErr bool
	}{
		{
			name: "ok",
			ts: Tags{
				Tag{
					Name:  "author",
					Value: "some author",
				},
				Tag{
					Name:  "album",
					Value: "some album",
				},
			},
		},
		{
			name: "err  duplicate tag",
			ts: Tags{
				Tag{
					Name:  "author",
					Value: "some author",
				},
				Tag{
					Name:  "author",
					Value: "some author",
				},
			},
			wantErr: true,
		},
		{
			name: "err  invalid tag",
			ts: Tags{
				Tag{
					Name:  "author",
					Value: "some author",
				},
				Tag{
					Name:  "album",
					Value: "",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ts.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Tags.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTag_Validate(t *testing.T) {
	type fields struct {
		Name  string
		Value string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Name:  "test",
				Value: "test",
			},
		},
		{
			name: "err  empty name",
			fields: fields{
				Name:  "",
				Value: "test",
			},
			wantErr: true,
		},
		{
			name: "err  empty value",
			fields: fields{
				Name:  "test",
				Value: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Tag{
				Name:  tt.fields.Name,
				Value: tt.fields.Value,
			}
			if err := tr.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Tag.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLyrics_Validate(t *testing.T) {
	tests := []struct {
		name    string
		l       Lyrics
		wantErr bool
	}{
		{
			name: "ok",
			l: Lyrics{
				{

					"some quote",
				},
			},
		},
		{
			name:    "err  empty lyrics",
			l:       Lyrics{},
			wantErr: true,
		},
		{
			name: "err  invalid verse",
			l: Lyrics{
				{
					"",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.l.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Lyrics.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVerse_Validate(t *testing.T) {
	tests := []struct {
		name    string
		v       Verse
		wantErr bool
	}{
		{
			name: "ok",
			v: Verse{
				"some quote",
			},
		},
		{
			name:    "err  empty verse",
			v:       make(Verse, 0),
			wantErr: true,
		},
		{
			name: "err  invalid quote",
			v: Verse{
				"some quote",
				"",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Verse.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQuote_Validate(t *testing.T) {
	tests := []struct {
		name    string
		q       Quote
		wantErr bool
	}{
		{
			name: "ok",
			q:    "some quote",
		},
		{
			name:    "err  empty quote",
			q:       "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.q.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Quote.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
