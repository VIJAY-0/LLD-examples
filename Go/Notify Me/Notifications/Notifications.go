package Notifications

type Observable interface {
	Add(Observer) error
	setValue(int) error
	NotifyAll() error
}

type Observer interface {
	Notify() error
	GetID() (int, error)
}
