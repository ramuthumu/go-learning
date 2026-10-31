package main

import "fmt"

// 1. Defining an Interface
// An interface is a contract of behaviors (method signatures).
// In Go, interfaces are satisfied IMPLICITLY (like Python duck typing, but type-safe!).
type Notifier interface {
	Send(message string) string
}

// 2. First Struct implementing Notifier
type EmailNotifier struct {
	Address string
}

func (e EmailNotifier) Send(message string) string {
	return fmt.Sprintf("[Email to %s]: %s", e.Address, message)
}

// 3. Second Struct implementing Notifier
type SMSNotifier struct {
	Phone string
}

func (s SMSNotifier) Send(message string) string {
	return fmt.Sprintf("[SMS to %s]: %s", s.Phone, message)
}

// 4. Function that accepts ANY Notifier
func AlertUser(n Notifier, msg string) {
	fmt.Println(n.Send(msg))
}

func main() {
	email := EmailNotifier{Address: "ramu@example.com"}
	sms := SMSNotifier{Phone: "555-0199"}

	// Both email and sms have a Send(string) string method,
	// so both automatically satisfy the Notifier interface!
	AlertUser(email, "Your server backup is complete.")
	AlertUser(sms, "Your verification code is 8492.")
}
