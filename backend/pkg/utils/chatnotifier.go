package utils

import (
	"fmt"
	"github.com/poncheska/terminal-http-chat/backend/pkg/models"
	"sync"
)

type ChatNotifier struct {
	mutex *sync.Mutex
	MemberChannels map[int64][]MemberChannel
}

type MemberChannel struct {
	MemberId int64
	Chan     chan models.MessageData
}

func NewChatNotifier() *ChatNotifier{
	return &ChatNotifier{
		mutex: &sync.Mutex{},
		MemberChannels: map[int64][]MemberChannel{},
	}
}

func (cn *ChatNotifier) Subscribe(memberId, chatId int64) (chan models.MessageData, error) {
	cn.mutex.Lock()
	defer cn.mutex.Unlock()

	for _, v := range cn.MemberChannels[chatId] {
		if v.MemberId == memberId {
			return nil, fmt.Errorf("user already subscribed this chat")
		}
	}
	ch := make(chan models.MessageData)
	mc := MemberChannel{memberId, ch}
	cn.MemberChannels[chatId] = append(cn.MemberChannels[chatId], mc)
	return ch, nil
}

func (cn *ChatNotifier) Unsubscribe(memberId, chatId int64) {
	cn.mutex.Lock()
	defer cn.mutex.Unlock()

	for i, v := range cn.MemberChannels {
		if v[chatId].MemberId == memberId {
			close(cn.MemberChannels[chatId][i].Chan)
			cn.MemberChannels[chatId] = append(cn.MemberChannels[chatId][:i], cn.MemberChannels[chatId][i+1:]...)
		}
	}
}

func (cn *ChatNotifier) Notify(chatId int64, m models.MessageData){
	cn.mutex.Lock()
	defer cn.mutex.Unlock()

	for _, v := range cn.MemberChannels[chatId]{
		v.Chan <- m
	}
}


