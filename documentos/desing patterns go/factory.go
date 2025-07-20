package main

import (
	"fmt"
)

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

func main() {
	notificationType := "email"

	if notificationType == "email" {
		notifier := EmailNotifier{}
		notifier.Send("Welcome!")
	} else if notificationType == "sms" {
		notifier := SMSNotifier{}
		notifier.Send("Welcome!")
	} else if notificationType == "push" {
		notifier := PushNotifier{}
		notifier.Send("Welcome!")
	}
}
