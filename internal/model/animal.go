package model

import "time"

type Animal struct {
	ID        uint
	Name      string
	Type      string
	Color     string
	CreatedAt time.Time
}
