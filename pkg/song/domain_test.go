package song

import (
	"reflect"
	"testing"
)

func TestSong_GetQuotes(t *testing.T) {
	type fields struct {
		Metadata Metadata
		Lyrics   Lyrics
	}
	tests := []struct {
		name   string
		fields fields
		want   []Quote
	}{
		{
			name: "ok",
			fields: fields{
				Lyrics: Lyrics{
					{
						Quotes: []Quote{
							{
								Phrase: "some phrase",
							},
							{
								Phrase: "another phrase",
							},
						},
					},
					{
						Quotes: []Quote{
							{
								Phrase: "one more phrase",
							},
						},
					},
				},
			},
			want: []Quote{
				{
					Phrase: "some phrase",
				},
				{
					Phrase: "another phrase",
				},
				{
					Phrase: "one more phrase",
				},
			},
		},
		{
			name: "empty",
			fields: fields{
				Lyrics: make(Lyrics, 0),
			},
			want: make([]Quote, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Song{
				Metadata: tt.fields.Metadata,
				Lyrics:   tt.fields.Lyrics,
			}
			if got := s.GetQuotes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Song.GetQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSong_GetRandomQuote(t *testing.T) {
	type fields struct {
		Metadata Metadata
		Lyrics   Lyrics
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Quote
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Lyrics: Lyrics{
					{
						Quotes: []Quote{
							{
								Phrase: "some phrase",
							},
							{
								Phrase: "some phrase",
							},
						},
					},
					{
						Quotes: []Quote{
							{
								Phrase: "some phrase",
							},
						},
					},
				},
			},
			want: &Quote{
				Phrase: "some phrase",
			},
		},
		{
			name: "empty",
			fields: fields{
				Lyrics: make(Lyrics, 0),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Song{
				Metadata: tt.fields.Metadata,
				Lyrics:   tt.fields.Lyrics,
			}
			got, err := s.GetRandomQuote()
			if (err != nil) != tt.wantErr {
				t.Errorf("Song.GetRandomQuote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Song.GetRandomQuote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSong_GetRandomVerse(t *testing.T) {
	type fields struct {
		Metadata Metadata
		Lyrics   Lyrics
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Verse
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Lyrics: Lyrics{
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
								Phrase: "some phrase",
							},
						},
					},
				},
			},
			want: &Verse{
				Quotes: []Quote{
					{
						Phrase: "some phrase",
					},
				},
			},
		},
		{
			name: "empty",
			fields: fields{
				Lyrics: make(Lyrics, 0),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Song{
				Metadata: tt.fields.Metadata,
				Lyrics:   tt.fields.Lyrics,
			}
			got, err := s.GetRandomVerse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Song.GetRandomVerse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Song.GetRandomVerse() = %v, want %v", got, tt.want)
			}
		})
	}
}
