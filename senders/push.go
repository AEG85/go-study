package senders

import "study/notifications"

type PushSender struct{}

func NewPushSender() *PushSender {
	return &PushSender{}
}

func (s PushSender) SendMessages(nl notifications.NotificationList) bool {
	if len(nl) > 0 {
		for _, v := range nl {
			s.sendPush(v)
		}
		return true
	}
	return false
}

func (s PushSender) sendPush(notifications.Notification) bool {
	return true
}
