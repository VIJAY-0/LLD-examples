package main

import (
	facade "SocialMedia/pkg/ActivityFacade"
	"SocialMedia/pkg/users"
	"fmt"
)

func main() {
	fmt.Println("Hello main")

	Activities := &facade.Facade{}
	user1 := &users.User{}
	Activities.UsersManager.AddUser(user1)
}
