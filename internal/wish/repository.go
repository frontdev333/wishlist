package wish

import "frontdev333/wishlist/database"

type Repository struct {
	db *database.DB
}

func (r *Repository) GetAll() []Wish {

	return nil
}

func (r *Repository) Get(slug string) Wish {

	return Wish{}
}

func (r *Repository) Store(payload *StorePayload) (Wish, error) {
	return Wish{}, nil
}

func (r *Repository) Update() (Wish, error) {
	return Wish{}, nil
}

func (r *Repository) Destroy(slug string) error {
	return nil
}
