package main

import (
	"github.com/samuelotter/i3ipc"
)

func Run() {

	channel := make(chan i3ipc.Event)

    socket, _ := i3ipc.GetIPCSocket()

    GetCurrent(socket)

    subscribeTo(i3ipc.I3WorkspaceEvent, channel)
	subscribeTo(i3ipc.I3WindowEvent, channel)

	for event := range channel {
        eventHandler(event, socket)
	}
}

func subscribeTo(eventType i3ipc.EventType, catcher chan i3ipc.Event) (err error) {
	channel, err := i3ipc.Subscribe(eventType)
	if err != nil {
		return
	}
	go func() {
		for e := range channel {
			catcher <- e
		}
	}()
	return
}
