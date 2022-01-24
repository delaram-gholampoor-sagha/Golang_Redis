package cache

import "github.com/Delaram-Gholampoor-Sagha/Golang_Redis/entity"

type PostCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}
