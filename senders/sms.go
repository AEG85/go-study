package senders

import "study/notifications"

type SmsSender struct{}

func NewSmsSender() *SmsSender {
	return &SmsSender{}
}

func (s SmsSender) SendMessages(nl notifications.NotificationList) bool {
	if len(nl) > 0 {
		for _, v := range nl {
			s.sendSms(v)
		}
		return true
	}
	return false
}

func (s SmsSender) sendSms(notifications.Notification) bool {
	return true
}
