package Notifications

import "sync"

type NotificationType struct {
}

type Notification struct {
	Type   NotificationType
	UserID int
	Body   string
}

type NotificationManager struct {
	notifications map[int][]*Notification
	mu            sync.RWMutex
}

func (nm *NotificationManager) AddNotification(userID int, notificationType NotificationType, content string) error {
	nm.mu.Lock()
	defer nm.mu.Unlock()

	return nil
}
