/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @since: 4/3/2023
  @desc: //TODO
**/

package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

type (
	ChatController struct {
	}
)

type SenderMsg struct {
	Receptor int64
	Msg      string
}

var Hub = make(map[int64]*websocket.Conn)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (c *ChatController) ChatWithUid(wsConn *websocket.Conn) {
	defer func(wsConn *websocket.Conn) {
		wsConn.CloseHandler()
	}(wsConn)
	for {
		msg := make([]byte, 512)
		pointMsg := new(SenderMsg)
		_, msg, err := wsConn.ReadMessage()
		if err != nil {

			return
		}
		err = json.Unmarshal(msg, pointMsg)
		if err != nil {
			err := wsConn.WriteMessage(websocket.TextMessage, []byte("Request Data Format Is Wrong"))
			if err != nil {
				return
			}
			continue
		}
		reWsConn, ok := Hub[pointMsg.Receptor]
		if !ok {
			err := wsConn.WriteMessage(websocket.TextMessage, []byte("This User Is Not Online"))
			if err != nil {
				return
			}
			continue
		}
		err = reWsConn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}

}
func (c *ChatController) WsHandle(gc *gin.Context) {
	userController := new(UserController)
	user, err := userController.GetUserInfo(gc)
	if err != nil {
		return
	}
	wsConn, err := upgrader.Upgrade(gc.Writer, gc.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	Hub[user.Id] = wsConn

	go c.ChatWithUid(wsConn)
}
func (c *ChatController) BroadcastMessage(gc *gin.Context) {

}
