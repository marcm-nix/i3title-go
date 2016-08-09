package main

import (
_    "fmt"
	"github.com/samuelotter/i3ipc"
)


//var ipcsocket i3ipc.IPCSocket


func run() {


	channel := make(chan i3ipc.Event)
	//subscribeTo(i3ipc.I3WorkspaceEvent, channel)
	subscribeTo(i3ipc.I3WindowEvent, channel)

//    ips, _ := i3ipc.GetIPCSocket()

	for event := range channel {
        eventHandler(event)
        handleEvent()
        //fmt.Printf("%s", a)
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
