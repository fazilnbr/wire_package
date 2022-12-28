package main

import "github.com/google/wire"

func initializeEvent() event {
	wire.Build(newMessage, newGreeter, newEvent)
	return event{}

}
