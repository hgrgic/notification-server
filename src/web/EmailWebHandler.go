package web

import (
	"co/iiq/i/notification-server/src/service"
	"fmt"
	"net/http"
	"strings"
)

func SendMailNotification(w http.ResponseWriter, r *http.Request) {

	urlQuery := r.URL.Query()
	subject := "Automated notification"
	message := urlQuery.Get("message")
	receiver := strings.Split(urlQuery.Get("receivers"), ",")
	bodyMessage := service.WritePlainEmail(receiver, subject, message)
	err := service.SendMail(receiver, subject, bodyMessage)

	if err != nil{
		fmt.Fprintf(w, "E-Mail Notificantion Failed %d", err)
		return
	}

	fmt.Fprintf(w, "E-Mail Notificantion Processed")
}
