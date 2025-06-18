package facade

import (
	"SocialMedia/pkg/Notifications"
	"SocialMedia/pkg/posts"
	"SocialMedia/pkg/users"
)

type Facade struct {
	NotificationsManager Notifications.NotificationManager
	UsersManager         users.UserManager
	PostsManager         posts.PostManager
}
