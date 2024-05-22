package models

import (
	"errors"
	"time"
)

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostCreateRequest struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Category string `json:"category" binding:"required"`
	Status   string `json:"status" binding:"required"`
}

type PostEditRequest struct {
	Title    string `json:"title,omitempty"`
	Content  string `json:"content,omitempty"`
	Category string `json:"category,omitempty"`
	Status   string `json:"status,omitempty"`
}

func (p PostCreateRequest) ValidateCreate() error {
	if len(p.Title) < 20 {
		return errors.New("judul harus memiliki minimal 20 karakter")
	}
	if len(p.Content) < 200 {
		return errors.New("konten harus memiliki minimal 200 karakter")
	}
	if len(p.Category) < 3 {
		return errors.New("kategori harus memiliki minimal 3 karakter")
	}
	if p.Status != "publish" && p.Status != "draft" && p.Status != "thrash" {
		return errors.New("status harus salah satu dari 'publish', 'draft', atau 'thrash'")
	}
	return nil
}

func (p PostEditRequest) ValidateUpdate() error {
	if p.Title != "" && len(p.Title) < 20 {
		return errors.New("judul harus memiliki minimal 20 karakter")
	}
	if p.Content != "" && len(p.Content) < 200 {
		return errors.New("konten harus memiliki minimal 200 karakter")
	}
	if p.Category != "" && len(p.Category) < 3 {
		return errors.New("kategori harus memiliki minimal 3 karakter")
	}
	if p.Status != "" && p.Status != "publish" && p.Status != "draft" && p.Status != "thrash" {
		return errors.New("status harus salah satu dari 'publish', 'draft', atau 'thrash'")
	}
	return nil
}
