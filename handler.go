package main

import (
    "fmt"
	"encoding/json"
	"github.com/samuelotter/i3ipc"
)


var format = "%s\n"


type ParsedWindowEvent struct {
	Container struct {
		Focused bool `json:"focused"`
		WindowProperties struct {
			Title string `json:"title"`
		} `json:"window_properties"`
	} `json:"container"`
}

func checkOutput(socket *i3ipc.IPCSocket) bool {

    if output == "all" {
        return true
    }


    i3workspaces, _ := socket.GetWorkspaces()

    for _, v := range i3workspaces {
        if v.Output == output && v.Focused == true {
            return true
        }
    }
    return false
}


func eventHandler(event i3ipc.Event, socket *i3ipc.IPCSocket) {

    //duplicated events, handled anyway
    if(event.Change == "new" || event.Change == "init" || event.Change == "close") {
        return
    }


    if checkOutput(socket) == false {
        return
    }

    if event.Change == "empty" {
        fmt.Printf(format, "")
        return
    }


    if event.Change == "focus" || event.Change == "title" {


        jsonString, _ := json.Marshal(event.Payload)
        var parsedWindowEvent = &ParsedWindowEvent{}
        json.Unmarshal([]byte(jsonString), &parsedWindowEvent)

        fmt.Printf(format, parsedWindowEvent.Container.WindowProperties.Title)

        return
    }



    fmt.Printf("change:%s\n\n", event.Change)


}


func iterateNodes(tree []i3ipc.I3Node, channel chan string) {

    for _, t := range tree {


        if t.Focused == true {
            channel <- t.Name
            break
        } else if len(t.Nodes) > 0 {

            go iterateNodes(t.Nodes, channel)
        }

    }
}




func GetCurrent(socket *i3ipc.IPCSocket) {

    tree, err := socket.GetTree()

    if err != nil {
        fmt.Printf(format, "")
        return
    }

    // if tree.Focused == true {
    //     //fmt.Println(tree.Name)
    //     return
    // }

    ch := make(chan string)

    for _, t := range tree.Nodes {

        if t.Layout == output || output == "all" {
            go iterateNodes(t.Nodes, ch)
        }
    }
    title := <-ch
    fmt.Printf(format, title)

    return
}
