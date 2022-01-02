package domain

import "gopatterns/repository/post"

// Repository must be implemented by all implementations of Post storage
type Repository interface {
	FindAll() ([]post.Post, error)
	Store(post post.Post) (post.Post, error)
	DeleteById(postId int64) error
}
