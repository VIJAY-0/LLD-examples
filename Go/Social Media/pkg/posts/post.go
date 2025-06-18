package posts

import (
	"sync"
)

type Post struct {
	ID       int
	UserID   int
	Content  string
	Likes    int
	Comments []*Comment
}

type Comment struct {
	text  string
	likes int
}

type PostManager struct {
	posts map[int]*Post
	mu    sync.RWMutex
}

func (pm *PostManager) GetPostbyID(postID int) (*Post, error) {
	return &Post{}, nil
}

func (pm *PostManager) AddPost(post *Post) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.posts[post.ID] = post
	return nil
}
