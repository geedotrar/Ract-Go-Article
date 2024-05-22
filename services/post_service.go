package services

import (
	"article-api/models"
	repositories "article-api/repostiories"
)

type PostService interface {
	GetAllPosts(offset, limit int) ([]models.Post, error)
	GetPostByID(id uint) (*models.Post, error)
	CreatePost(post *models.Post) error

	UpdatePost(id uint, post *models.Post) error

	DeletePost(id uint) error
}

type postService struct {
	postRepo repositories.PostRepository
}

func NewPostService(repo repositories.PostRepository) PostService {
	return &postService{repo}
}

func (s *postService) GetAllPosts(offset, limit int) ([]models.Post, error) {
	return s.postRepo.GetAll(offset, limit)
}

func (s *postService) GetPostByID(id uint) (*models.Post, error) {
	return s.postRepo.GetByID(id)
}

func (s *postService) CreatePost(post *models.Post) error {
	return s.postRepo.Create(post)
}

func (s *postService) UpdatePost(id uint, post *models.Post) error {
	return s.postRepo.UpdatePost(id, post)
}

func (s *postService) DeletePost(id uint) error {
	return s.postRepo.Delete(id)
}
