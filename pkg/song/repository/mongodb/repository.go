package mongodb

type collection interface {
}

type Repository struct {
	c collection
}

func NewRepository(c collection) (*Repository, error) {
	return &Repository{
		c: c,
	}, nil
}
