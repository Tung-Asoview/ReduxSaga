package services

import (
	"Go-Chi/Model"
)

type PostService interface {
	AllPosts() ([]Model.Post, error)
	//DetailPost(id int) (*Model.Post, error)
	//CreatePost(u *Model.Post) (error)
	//UpdatePost(id int) (error)
	//DeletePost(id int) (error)
}

//func AllPosts() ([]Model.Post, error) {
//	note, err := repository.AllPosts()
//	return note, err
//}
//
//func DetailPost(id string) (Model.Post) {
//	note := repository.DetailPost(id)
//	return note
//}
//
//func CreatePost(p Model.Post) {
//	repository.CreatePost(p)
//}