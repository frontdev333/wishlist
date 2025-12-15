package wish

import "frontdev333/wishlist/database"

type Wish struct {
	ID                uint    `json:"id"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	ForeignLink       string  `json:"foreign_link"`
	Price             float64 `json:"price"`
	IsMultipleBooking bool    `json:"is_multiple_booking"`
	Slug              string  `json:"slug"`
	OwnerID           uint    `json:"owner_id"`
	db                *database.DB
}

func (w *Wish) FindBySlug(slug string) Wish {
	return Wish{}
}
