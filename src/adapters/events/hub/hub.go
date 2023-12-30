package hub

import (
	"github.com/leandro-lugaresi/hub"
	eventDomain "godeck/src/domain/events"
	"time"
)

type Hub struct {
	hub *hub.Hub

	later map[eventDomain.Event]time.Time
}

func (h *Hub) Publish(topic eventDomain.Event, data eventDomain.EventData) {
	h.hub.Publish(hub.Message{
		Name:   string(topic),
		Fields: hub.Fields(data),
	})
}

func (h *Hub) PublishLater(topic eventDomain.Event, data eventDomain.EventData, delay time.Duration) {
	time.AfterFunc(delay, func() {
		if h.later[topic].IsZero() {
			return
		}
		h.Publish(topic, data)
	})
	h.later[topic] = time.Now().Add(delay)
}

func (h *Hub) CancelPublishLater(topic eventDomain.Event) {
	if h.later[topic].IsZero() {
		return
	}

	h.later[topic] = time.Time{}
}

func (h *Hub) Subscribe(topic eventDomain.Event, notify eventDomain.EventListener) {
	sub := h.hub.Subscribe(10, string(topic))

	go func(s hub.Subscription) {
		for msg := range s.Receiver {
			notify(topic, eventDomain.EventData(msg.Fields))
		}
	}(sub)
}

func NewHub() *Hub {
	return &Hub{
		hub: hub.New(),
		later: map[eventDomain.Event]time.Time{
			eventDomain.DeviceWillSleepEvent: time.Time{},
		},
	}
}