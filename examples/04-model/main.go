package main

import (
	"errors"
	"fmt"
)

type Ticket struct {
	Title  string
	Closed bool
}

func (t *Ticket) Close() error {
	if t.Closed {
		return errors.New("ticket already closed")
	}
	t.Closed = true
	return nil
}

func main() {
	ticket := Ticket{Title: "Fix the export"}
	copyOf := ticket
	if err := ticket.Close(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ticket.Closed, copyOf.Closed)
	fmt.Println(ticket.Close())
}
