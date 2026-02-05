package senders

import "study/notifications"

type Sender interface {
	SendMessages(notifications.NotificationList) bool
}
