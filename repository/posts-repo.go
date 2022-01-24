package repository

import "github.com/Delaram-Gholampoor-Sagha/Golang_Redis/entity"



type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindByID(id string) (*entity.Post, error)
	Delete(post *entity.Post) error
}
