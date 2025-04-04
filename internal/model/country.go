package model

import (
	"time"

	"github.com/google/uuid"
)

type Country struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Capital   string    `json:"capital"`
	Continent string    `json:"continent"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
