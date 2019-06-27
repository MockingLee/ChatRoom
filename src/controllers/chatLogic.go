package controllers

import (
	"../model"
	"container/list"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
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
	subscribers = list.New()
	waitingList = list.New()
)

func doChat(){
	for{
		select {
		case u := <- online_user:
			fmt.Println("user in : " , u.Name ,";WebSocket:", u.Conn != nil)
			beego.Info("user in :", u.Name, ";WebSocket:", u.Conn != nil)
		case u := <- offline_user:
			beego.Info("user off : " , u)
		case event := <- publish:
			// Notify waiting list.
			for ch := waitingList.Back(); ch != nil; ch = ch.Prev() {
				ch.Value.(chan bool) <- true
				waitingList.Remove(ch)
			}

			broadcastWebSocket(event)
			model.NewArchive(event)

			if event.Type == model.EVENT_MESSAGE {
				beego.Info("Message from", event.User, ";Content:", event.Content)
			}
			

		}
	}
}

func init(){
	go doChat()
}

