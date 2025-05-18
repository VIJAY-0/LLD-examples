package main

import (
	"Notifyme/Notifications"
	"fmt"
)

func main() {
	fmt.Println("Helloddd")

	var a Notifications.EmailNotification

	user1 := Notifications.EmailObserver{ID: 0, Email: "1123@example.com"}
	user2 := Notifications.EmailObserver{ID: 1, Email: "1223@example.com"}
	user3 := Notifications.EmailObserver{ID: 2, Email: "1233@example.com"}

	a.Add(user1)
	a.Add(user2)
	a.Add(user3)

	a.SetValue(30)

}
