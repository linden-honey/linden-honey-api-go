package song

// Entity is a main domain object.
type Entity struct {
	Metadata
	Lyrics Lyrics `json:"lyrics"`
}

// Metadata is a domain helper object.
type Metadata struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Tags  Tags   `json:"tags,omitempty"`
}

// Tags is a domain object.
type Tags []Tag

// Tag is a domain object.
type Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Lyrics is a domain object.
type Lyrics []Verse

// Verse is a domain object.
type Verse []Quote

// Quote is a domain object.
type Quote string
