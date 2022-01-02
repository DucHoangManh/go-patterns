package memstorage

import (
	"fmt"
	_ "gopatterns/repository/domain"
	"gopatterns/repository/post"
)

//in-memory implementation
type PostStorage struct {
	posts     map[int64]string
	highestID int64
}

func (p *PostStorage) FindAll() ([]post.Post, error) {
	result := make([]post.Post, 0)
	for id, title := range p.posts {
		result = append(result, post.Post{
			ID:    id,
			Title: title,
		})
	}
	return result, nil
}

func (p *PostStorage) Store(post post.Post) (post.Post, error) {
	if p.posts == nil {
		p.posts = make(map[int64]string)
	}
	if post.ID <= 0 {
		p.highestID++
		post.ID = p.highestID
	} else {
		if _, exists := p.posts[post.ID]; !exists {
			return post, fmt.Errorf("post already exist")
		}
	}
	p.posts[post.ID] = post.Title
	return post, nil

}

func (p *PostStorage) DeleteById(postId int64) error {
	delete(p.posts, postId)
	return nil
}
