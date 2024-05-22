package repositories

import (
	"article-api/models"
	"time"

	"gorm.io/gorm"
)

type PostRepository interface {
	GetAll(offset, limit int) ([]models.Post, error)
	GetByID(id uint) (*models.Post, error)
	Create(post *models.Post) error
	UpdatePost(id uint, post *models.Post) error
	Delete(id uint) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

func (r *postRepository) GetAll(offset, limit int) ([]models.Post, error) {
	var posts []models.Post
	if err := r.db.Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) GetByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := r.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) Create(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) Delete(id uint) error {
	return r.db.Delete(&models.Post{}, id).Error
}

func (r *postRepository) UpdatePost(id uint, post *models.Post) error {
	return r.db.Model(&models.Post{}).Where("id = ?", id).Updates(models.Post{
		Title:     post.Title,
		Content:   post.Content,
		Category:  post.Category,
		Status:    post.Status,
		UpdatedAt: time.Now(),
	}).Error
}
