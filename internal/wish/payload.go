package wish

type StorePayload struct {
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	ForeignLink       string  `json:"foreign_link"`
	Price             float64 `json:"price"`
	IsMultipleBooking bool    `json:"is_multiple_booking"`
	Slug              string  `json:"slug"`
	OwnerID           uint    `json:"owner_id"`
}
