package main

import (
	"gopatterns/repository/memstorage"
	"gopatterns/repository/post"
	"log"
)

func main() {
	postRepo := memstorage.PostStorage{}
	//may change to other implementation in the future
	newPost, err := postRepo.Store(post.Post{
		Title: "Eagles fly",
	})
	if err != nil {
		log.Println("can not create post", err)
	} else {
		log.Printf("created post with id %d", newPost.ID)
	}
	posts, err := postRepo.FindAll()
	if err != nil {
		log.Println("can not fetch posts", err)
	} else {
		log.Println(posts)
	}
}
