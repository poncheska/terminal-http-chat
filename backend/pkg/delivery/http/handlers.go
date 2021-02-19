package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/poncheska/terminal-http-chat/backend/pkg/models"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (h Handler) Chat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chatId, err:= strconv.ParseInt(vars["id"],10,64)
	if err != nil {
		log.Println("chat: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}
	userId, err:= strconv.ParseInt(r.Header.Get(userIdHeader),10,64)
	if err != nil {
		log.Println("chat: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}
	msgs, err := h.store.View.GetAllMessageData(chatId)
	if err != nil {
		log.Println("chat: store error: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Println("chat: socket upgrader error: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}

	for _, v := range msgs{
		err := conn.WriteJSON(v)
		if err != nil {
			WriteErrorResponse(w, err)
			return
		}
	}

	go h.ChatStream(conn, userId, chatId)
}

func (h Handler) ChatNotification(conn *websocket.Conn, userId, chatId int64){
	ch, err := h.chatNotifier.Subscribe(userId, chatId)
	if err != nil{
		log.Printf(err.Error())
		return
	}
	defer h.chatNotifier.Unsubscribe(userId, chatId)
	for{
		msgData := <-ch
		err := conn.WriteJSON(msgData)
		if err != nil{
			log.Printf(err.Error())
			return
		}
	}
}

func (h Handler) ChatStream(conn *websocket.Conn, userId, chatId int64){
	defer conn.Close()
	for{
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf(err.Error())
			return
		}

		var parsedMsg *models.Message
		err = json.Unmarshal(msg, parsedMsg)
		if err != nil {
			log.Printf(err.Error())
			return
		}
		parsedMsg.Id = userId
		parsedMsg.Date = time.Now()

		msgId, err := h.store.Message.Create(*parsedMsg)
		if err != nil {
			log.Printf(err.Error())
			return
		}

		msgData, err := h.store.View.GetMessageDataById(msgId)
		if err != nil {
			log.Printf(err.Error())
			return
		}

		h.chatNotifier.Notify(chatId, msgData)
	}
}

func (h Handler) Chats(w http.ResponseWriter, r *http.Request) {
	chats, err := h.store.Chat.GetAll()
	if err != nil {
		log.Println("chats: store error: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(chats)
	if err != nil{
		log.Println("chats: response encode error: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (h Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	data := &AuthData{}
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		log.Println("signin: request decode error: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}
	if err := data.CheckNotEmpty(); err != nil{
		log.Println("signin: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}

	user, err := h.store.User.GetByCredentials(data.Username, data.Password)
	if err != nil {
		log.Println("signin: store error: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}

	token, err := h.tokenService.CreateToken(user.Id)
	if err != nil {
		log.Println("signin: token create error: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
	if err != nil {
		log.Println("signin: response encode error: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (h Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	data := &AuthData{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		log.Println("signup: request decode error: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}
	if err := data.CheckNotEmpty(); err != nil{
		log.Println("signup: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}

	_, err = h.store.User.Create(data.Username, data.Password)
	if err != nil {
		log.Println("signup: store error: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}

	//token, err := h.tokenService.CreateToken(id)
	//if err != nil {
	//	WriteErrorResponse(w, err)
	//	return
	//}
	//
	//err = json.NewEncoder(w).Encode(map[string]string{
	//	"token": token,
	//})
	//if err != nil {
	//	WriteErrorResponse(w, err)
	//	return
	//}
	//w.Header().Set("Content-Type", "application/json")
}

func (h Handler) CreateChat(w http.ResponseWriter, r *http.Request) {
	userId, err:= strconv.ParseInt(r.Header.Get(userIdHeader),10,64)
	if err != nil {
		log.Println("create_chat: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}

	var ccr CreateChatReq
	if err := json.NewDecoder(r.Body).Decode(&ccr); err != nil {
		log.Println("create_chat: request decode error: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}
	if err := ccr.CheckNotEmpty(); err != nil{
		log.Println("create_chat: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}

	chat := models.Chat{
		Name: ccr.ChatName,
		AdminId: userId,
	}
	_, err = h.store.Chat.Create(chat)
	if err != nil {
		log.Println("create_chat: store error: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}
}

func (h Handler) DeleteChat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chatId, err:= strconv.ParseInt(vars["id"],10,64)
	if err != nil {
		log.Println("delete_chat: " + err.Error())
		return
	}
	userId, err:= strconv.ParseInt(r.Header.Get(userIdHeader),10,64)
	if err != nil {
		log.Println("delete_chat: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}
	err = h.store.Chat.Delete(chatId, userId)
	if err != nil {
		log.Println("delete_chat: store error: " + err.Error())
		WriteErrorResponse(w, err)
		return
	}
}