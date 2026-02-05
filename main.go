package main

import (
	"strconv"
	"study/notifications"
	"study/senders"

	"github.com/k0kubun/pp/v3"
)

func main() {
	notifireList := notifications.NewNotificationList(100)

	for i := range 100 {
		notifire := notifications.NewNotification(strconv.Itoa(i), "Номер сообщения "+strconv.Itoa(i))
		notifireList.AddNotification(notifire)
	}
	pp.Println("Сообщения", notifireList)
	sender := senders.NewSmsSender()
	ok := sender.SendMessages(*notifireList)
	if ok {
		pp.Println("Все успешно отправлено!")
	} else {
		pp.Println("При отправки что-то пошло не так!")
	}

}
