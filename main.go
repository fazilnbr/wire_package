package main

import "fmt"

// In main function we call the funtions with dependancies

func main() {
	fmt.Print("\n\n--welcome to fazilnbr's wire_package repostiory--\n\n\n\n")

	e := initializeEvent()

	e.start()

	fmt.Println("\nProgram Ended..")

}

// Initialize the struct, struct methord and it's constructors with depentancies

// First message struct, struct methord and it's constructors

type message struct {
	value string
}

func newMessage(msg string) message {
	return message{value: msg}
}

func (m message) string() string {
	return fmt.Sprintf("[OFFICIAL MESSAGE] : %s", m.value)
}

// Second greeter struct, struct methord and it's constructors
// It depend on the message struct

type greeter struct {
	message message
}

func newGreeter(m message) greeter {
	return greeter{message: m}
}

func (g greeter) greet() string {
	return g.message.string()
}

// third event struct, struct methord and it's constructors
// It depend on the greeter struct and greeter struct depend on the message struct

type event struct {
	greeter greeter
}

func newEvent(g greeter) event {
	return event{greeter: g}
}

func (e event) start() {
	fmt.Printf("Output : %s\n\n\n", e.greeter.greet())
}
