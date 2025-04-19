package models

import (
	"time"

	"github.com/rs/xid"
)

type Country struct {
	ID        xid.ID    `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Capital   string    `json:"capital"`
	Continent string    `json:"continent"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
