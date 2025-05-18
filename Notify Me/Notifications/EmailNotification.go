package Notifications

import (
	"fmt"
	"time"
)

type Attachments struct {
}

type EmailNotification struct {
	value       int
	Subject     string
	Body        string
	Attachments Attachments

	Subscribers []EmailObserver
}

type EmailObserver struct {
	ID    int
	Email string
}

func (e *EmailNotification) Add(eo EmailObserver) error {
	e.Subscribers = append(e.Subscribers, eo)
	return nil
}

func (e *EmailNotification) SetValue(value int) error {
	e.value = value
	e.NotifyAll()
	return nil
}
func (e *EmailNotification) NotifyAll() error {

	for _, obs := range e.Subscribers {
		obs.Notify()
	}

	return nil
}

func (eo *EmailObserver) Notify() error {
	fmt.Printf("sending mail to  [%s] \n", eo.Email)
	time.Sleep(1500 * time.Millisecond)
	fmt.Printf("sent mail to  [%s] \n", eo.Email)
	return nil
}

func (eo *EmailObserver) GetID() (int, error) {

	return eo.ID, nil
}
