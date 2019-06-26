package controllers

import (
	"fmt"
	"github.com/gorilla/websocket"
	"../model"
	"time"
)

func newEvent(ep model.EventType, user, msg string) model.Event {
	return model.Event{ep, user, int(time.Now().Unix()), msg}
}

func Join(user string, ws *websocket.Conn) {
	online_user <- Subscriber{Name: user, Conn: ws}
}

func Leave(user string) {
	offline_user <- user
}
type Subscriber struct {
	Name string
	Conn *websocket.Conn
}

var (
	online_user = make(chan Subscriber , 10)
	offline_user = make(chan string , 10)
	publish = make(chan model.Event, 10)

)

func doChat(){
	for{
		select {
		case u := <-online_user:
			fmt.Println("user in : " , u.Name ,";WebSocket:", u.Conn != nil)
		//case event := <- publish:
		}
	}
}

