package service

import (
	"bytes"
	"co/iiq/i/notification-server/src/model"
	"co/iiq/i/notification-server/src/util"
	"fmt"
	"log"
	"mime/quotedprintable"
	"net/smtp"
	"strings"
)

var (
	props, propErr = util.ReadPropertiesFile("resources/connection.properties")
	sender = model.NewSender(props["emailuser"], props["emailpass"])
	SMTPServer = props["SMTPServer"]
	SMTPPort = props["SMTPPort"]
)


func SendMail(Dest []string, Subject, bodyMessage string) error {
	if propErr != nil {
		log.Println("Email properties not loaded correctly.", propErr)
		return propErr
	}

	msg := "From: " + sender.User + "\n" +
		"To: " + strings.Join(Dest, ",") + "\n" +
		"Subject: " + Subject + "\n" + bodyMessage

	err := smtp.SendMail(SMTPServer+":"+SMTPPort,
		smtp.PlainAuth("", sender.User, sender.Password, SMTPServer),
		sender.User, Dest, []byte(msg))

	if err != nil {
		fmt.Printf("smtp error: %s", err)
		return err
	}

	return nil
}



func writeEmail(dest []string, contentType, subject, bodyMessage string) string {
	header := make(map[string]string)
	header["From"] = sender.User

	receipient := ""

	for _, user := range dest {
		receipient = receipient + user
	}

	header["To"] = receipient
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(bodyMessage))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}


func WriteHTMLEmail(dest []string, subject, bodyMessage string) string {
	return writeEmail(dest, "text/html", subject, bodyMessage)
}

func WritePlainEmail(dest []string, subject, bodyMessage string) string {
	return writeEmail(dest, "text/plain", subject, bodyMessage)
}
