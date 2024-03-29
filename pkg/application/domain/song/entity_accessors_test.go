package song

import (
	"reflect"
	"testing"
)

func TestLyrics_GetRandomVerse(t *testing.T) {
	tests := []struct {
		name    string
		l       Lyrics
		want    Verse
		wantErr bool
	}{
		{
			name: "ok",
			l: Lyrics{
				{
					"some quote",
				},
				{
					"some quote",
				},
			},
			want: Verse{
				"some quote",
			},
		},
		{
			name:    "err  empty lyrics",
			l:       Lyrics{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.GetRandomVerse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Lyrics.GetRandomVerse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lyrics.GetRandomVerse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLyrics_GetRandomQuote(t *testing.T) {
	tests := []struct {
		name    string
		l       Lyrics
		want    Quote
		wantErr bool
	}{
		{
			name: "ok",
			l: Lyrics{
				{

					"some quote",
					"some quote",
				},
				{

					"some quote",
				},
			},
			want: "some quote",
		},
		{
			name:    "empty",
			l:       Lyrics{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.GetRandomQuote()
			if (err != nil) != tt.wantErr {
				t.Errorf("Lyrics.GetRandomQuote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lyrics.GetRandomQuote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLyrics_GetQuotes(t *testing.T) {
	tests := []struct {
		name string
		l    Lyrics
		want []Quote
	}{
		{
			name: "ok",
			l: Lyrics{
				{

					"some quote",
					"another quote",
				},
				{

					"one more quote",
				},
			},
			want: []Quote{
				"some quote",
				"another quote",
				"one more quote",
			},
		},
		{
			name: "empty",
			l:    Lyrics{},
			want: make([]Quote, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.GetQuotes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lyrics.GetQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}
