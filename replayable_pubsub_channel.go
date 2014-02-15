package replayable_pubsub_channel

import (
	"errors"
	"container/list"
)

type Replayable_pubsub_channel struct
{
	subscribers list.List
}

func (c Replayable_pubsub_channel) lazyInit() {
	//if c.subscribers == nil {
	//	c.subscribers = list.New()
	//}	
}

func (c Replayable_pubsub_channel) Publish(message string) (err error) {
	c.lazyInit()
	for e := c.subscribers.Front(); e != nil; e.Next() {
		e.Value.(chan string) <- message
	}	
	return
}

func (c Replayable_pubsub_channel) Subscribe(receiver chan string) (err error) {
	c.lazyInit()
	c.subscribers.PushBack(receiver)
	return
}

func (c Replayable_pubsub_channel) Unsubscribe(receiver chan string) (err error) {
	return errors.New("Not implemented")
}