package payload

import "time"

type CreateMoviePayload struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Rating      float64 `json:"rating" validate:"required"`
	Image       string  `json:"image"`
}

type UpdateMoviePayload struct {
	Id          int64   `json:"-" validate:"required"`
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Rating      float64 `json:"rating" validate:"required"`
	Image       string  `json:"image"`
}

type MoviePayload struct {
	Id          int64      `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Rating      float64    `json:"rating"`
	Image       string     `json:"image"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
