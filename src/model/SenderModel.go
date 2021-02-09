package model

type Sender struct {
	User     string
	Password string
}

func NewSender(Username, Password string) Sender {
	return Sender{Username, Password}
}
