package entity

import "time"

type Tag struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Tags []Tag

type PubicTag struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (t *Tag) PublicTag() *PubicTag {
	return &PubicTag{
		Id:          t.Id,
		Name:        t.Name,
		Description: t.Description,
	}
}
