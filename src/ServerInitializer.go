package main

import (
	"co/iiq/i/notification-server/src/web"
	"fmt"
	"log"
	"net/http"
)

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Notificantion Server Operational")
}


func main() {
	http.HandleFunc("/", status)
	http.HandleFunc("/email", web.SendMailNotification)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
