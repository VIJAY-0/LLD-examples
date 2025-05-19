package users

import (
	"SocialMedia/pkg/posts"
	"sync"
)

type User struct {
	ID             string
	Username       string
	Name           string
	Email          string
	Password       string
	DisplayPicture *string
	Bio            string
	Friends        map[int]*User
	Posts          []*posts.Post
}

type UserManager struct {
	users map[int]*User
	mu    sync.RWMutex
}

func (um *UserManager) AddUser(user *User) error {
	return nil
}

func (um *UserManager) GetUserbyID(uid int) (*User, error) {

	return &User{}, nil
}

func (um *UserManager) AddFriend(requesterID int, recieverID int) error {
	return nil
}
