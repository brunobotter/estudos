package main

import (
	"errors"
	"fmt"
)

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (n EmailNotifier) Send(message string) {
	fmt.Println("Sending EMAIL:", message)
}

type SMSNotifier struct{}

func (n SMSNotifier) Send(message string) {
	fmt.Println("Sending SMS:", message)
}

type PushNotifier struct{}

func (n PushNotifier) Send(message string) {
	fmt.Println("Sending PUSH:", message)
}

func NotifierFactory(t string) (Notifier, error) {
	switch t {
	case "email":
		return EmailNotifier{}, nil
	case "sms":
		return SMSNotifier{}, nil
	case "push":
		return PushNotifier{}, nil
	default:
		return nil, errors.New("unknown notifier type")
	}
}

func main() {
	notificationType := "push"

	notifier, err := NotifierFactory(notificationType)
	if err != nil {
		panic(err)
	}

	notifier.Send("Welcome!")
}
