package entity

import (
	"errors"
	"time"
)

type Article struct {
	Id        int       `json:"id" gorm:"primary_key;autoIncrement"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorId  uint64    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Article) VaildContentInfo() (bool, error) {
	if len(a.Title) > 80 {
		return false, errors.New("title must less than 80 character")
	}
	if a.AuthorId == 0 {
		return false, errors.New("authorId must not be empty")
	}
	return true, nil
}

func (a *Article) PrepareCreate() {
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
}
