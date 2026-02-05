package senders

import "study/notifications"

type MailSender struct{}

func NewMailSender() *MailSender {
	return &MailSender{}
}

func (s MailSender) SendMessages(nl notifications.NotificationList) bool {
	if len(nl) > 0 {
		for _, v := range nl {
			s.sendMail(v)
		}
		return true
	}
	return false
}

func (s MailSender) sendMail(notifications.Notification) bool {
	return true
}
