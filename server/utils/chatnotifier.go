package utils

import (
	"fmt"
	"terminal-http-chat/server/models"
)

type ChatNotifier struct {
	MemberChannels []MemberChannel
}

type MemberChannel struct {
	MemberId int64
	Chan     chan models.Message
}

func (cn *ChatNotifier) Subscribe(memberId int64) (chan models.Message, error) {
	for _, v := range cn.MemberChannels {
		if v.MemberId == memberId {
			return nil, fmt.Errorf("user already subscribed this chat")
		}
	}
	ch := make(chan models.Message)
	mc := MemberChannel{memberId, ch}
	cn.MemberChannels = append(cn.MemberChannels, mc)
	return ch, nil
}

func (cn *ChatNotifier) Unsubscribe(memberId int64) {
	for i, v := range cn.MemberChannels {
		if v.MemberId == memberId {
			cn.MemberChannels = append(cn.MemberChannels[:i], cn.MemberChannels[i+1:]...)
		}
	}
}

func (cn *ChatNotifier) Notify(m models.Message){
	for _, v := range cn.MemberChannels{
		v.Chan <- m
	}
}


