package controllers

import (
	"ChatRoom/src/model"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
)

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	baseController
}

// Get method handles GET requests for WebSocketController.
func (this *WebSocketController) Get() {
	// Safe check.
	uname := this.GetString("username")
	if len(uname) == 0 {
		this.Data["json"] = "{\"success\" : \"false\" , \"msg\" : \"username can not be null\"}"
		this.ServeJSON()
		return
	}

	for u := subscribers.Front() ; u != nil ; u = u.Next() {
		if uname == u.Value.(Subscriber).Name {
			this.Data["json"] = "{\"success\" : \"false\" , \"msg\" : \"username exits\"}"
			this.ServeJSON()
			return
		}
	}


	this.Data["json"] = "{\"success\" : \"true\"}"
	this.ServeJSON()

}

// Join method handles WebSocket requests for WebSocketController.
func (this *WebSocketController) Join() {
	uname := this.GetString("username")
	if len(uname) == 0 {
		this.Data["json"] = "{\"success\" : \"false\" , \"msg\" : \"username can not be null\"}"
		this.ServeJSON()
		return
	}

	for u := subscribers.Front() ; u != nil ; u = u.Next() {
		if uname == u.Value.(Subscriber).Name {
			this.Data["json"] = "{\"success\" : \"false\" , \"msg\" : \"username exits\"}"
			this.ServeJSON()
			return
		}
	}

	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	//ws.WriteJSON()
	// Join chat room.
	Join(uname, ws)
	defer Leave(uname)

	// Message receive loop.
	for {
		_, p, err := ws.ReadMessage()

		if err != nil {
			return
		}
		beego.Info(uname , " says " , string(p))
		publish <- newEvent(model.EVENT_MESSAGE, uname, string(p))
	}
}

// broadcastWebSocket broadcasts messages to WebSocket users.
func broadcastWebSocket(event model.Event) {
	data, err := json.Marshal(event)
	fmt.Print(data)
	if err != nil {
		beego.Error("Fail to marshal event:", err)
		return
	}
	fmt.Println(subscribers.Len())
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		// Immediately send event to WebSocket users.
		ws := sub.Value.(Subscriber).Conn
		//fmt.Print(ws)
		if ws != nil {
			//beego.Info("send msg to ",sub)
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				// User disconnected.
				unsubscriber <- sub.Value.(Subscriber).Name
			}
		}
	}
}