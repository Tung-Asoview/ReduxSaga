package repository

import (
	models "Go-Chi/Model"
)

type PostRepo interface {
	AllPosts() ([]*models.Post, error)
	DetailPost(id int) (*models.Post, error)
	CreatePost(u *models.Post) (error)
	UpdatePost(id int) (error)
	DeletePost(id int) (error)
}
