package controllers

import (
	"ChatRoom/src/model"
	"container/list"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"time"
)

func newEvent(ep model.EventType, user, msg string) model.Event {
	return model.Event{ep, user, int(time.Now().Unix()), msg}
}

func Join(user string, ws *websocket.Conn) {
	subscriber <- Subscriber{Name: user, Conn: ws}
}

func Leave(user string) {
	unsubscriber <- user
}
type Subscriber struct {
	Name string
	Conn *websocket.Conn
}

var (
	subscriber = make(chan Subscriber , 10) //user in channel
	unsubscriber = make(chan string , 10)  //user out channel
	publish = make(chan model.Event, 10) //new event channel
	subscribers = list.New() //online users list
	waitingList = list.New() //offline users list
)

func doChat(){
	for{
		select {
		case u := <- subscriber:
			//publish <- newEvent(model.EVENT_JOIN, u.Name, "")
			//fmt.Println("user in : " , u.Name ,";WebSocket:", u.Conn != nil)
			subscribers.PushBack(u)
			beego.Info("user in :", u.Name, ";WebSocket:", u.Conn != nil)
			publish <- newEvent(model.EVENT_JOIN , u.Name , "")
		case u := <- unsubscriber:
			for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).Name == u {
					subscribers.Remove(sub)
					// Clone connection.
					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
						beego.Error("WebSocket closed:", u)
					}
					publish <- newEvent(model.EVENT_LEAVE, u, "") // Publish a LEAVE event.
					break
				}
			}
		case event := <- publish:
			// Notify waiting list.
			/*for ch := waitingList.Back(); ch != nil; ch = ch.Prev() {
				ch.Value.(chan bool) <- true
				waitingList.Remove(ch)
			}*/

			broadcastWebSocket(event)
			//model.NewArchive(event)

			if event.Type == model.EVENT_MESSAGE {
				beego.Info("Message from", event.User, ";Content:", event.Content)
			}
			

		}
	}
}

func init(){
	go doChat()
}

