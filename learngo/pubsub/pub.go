package pubsub

import (
	"sync"
	"time"
)

type Publisher struct {
	m          sync.RWMutex
	buffer     int
	timeout    time.Duration
	subscribes map[subscriber]topicFunc
}

func NewPubulisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:     buffer,
		timeout:    publishTimeout,
		subscribes: make(map[subscriber]topicFunc),
	}
}
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribes[ch] = topic
	p.m.Unlock()
	return ch
}
