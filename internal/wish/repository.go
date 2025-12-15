package wish

type Repository struct {
	model *Wish
}

func (r *Repository) GetAll() []Wish {

	return nil
}

func (r *Repository) Get(slug string) Wish {

	return Wish{}
}

func (r *Repository) Store() (Wish, error) {
	return Wish{}, nil
}

func (r *Repository) Update() (Wish, error) {
	return Wish{}, nil
}

func (r *Repository) Destroy(slug string) error {
	return nil
}
