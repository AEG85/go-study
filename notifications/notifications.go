package notifications

import "github.com/k0kubun/pp/v3"

type Notification struct {
	Title   string
	Message string
}

// Определите тип для списка уведомлений
type NotificationList []Notification

func NewNotification(title string, message string) Notification {
	return Notification{
		Title:   title,
		Message: message,
	}
}

func NewNotificationList(capacity int) *NotificationList {
	list := make(NotificationList, 0, capacity)
	return &list
}

func (n *NotificationList) AddNotification(notifire Notification) bool {
	if notifire.Title == "" {
		pp.Println("Не указан заголовок сообщения!")
		return false
	}

	if notifire.Message == "" {
		pp.Println("Не указано описание сообщения!")
		return false
	}
	*n = append(*n, notifire)
	return true
}
