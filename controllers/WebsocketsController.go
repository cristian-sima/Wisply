package controllers

import ws "github.com/cristian-sima/Wisply/models/ws"

var hub *ws.Hub

func init() {
	hub = ws.CreateHub()
	go hub.Run()
}

// WebsocketsController It manages the operations for repository (list, delete, add)
type WebsocketsController struct {
	AdminController
}

// InitWebsocketConnection Initiats the websocket connection
func (controller *WebsocketsController) InitWebsocketConnection() {
	controller.TplNames = "site/harvest/init.tpl"
	connection := hub.CreateConnection(controller.Ctx.ResponseWriter, controller.Ctx.Request, controller)
	hub.Register <- connection
	go connection.WritePump()
	connection.ReadPump()
}
