package replayable_pubsub_channel

import (
	"container/list"
)

type Replayable_pubsub_channel struct
{
	Subscribers *list.List
}

func (c *Replayable_pubsub_channel) lazyInit() {
	if c.Subscribers == nil {
		c.Subscribers = list.New()
	}	
}

func (c *Replayable_pubsub_channel) Publish(message string) {
	c.lazyInit()	
		
	for e := c.Subscribers.Front(); e != nil; e = e.Next() {				
		e.Value.(chan string) <- message	
	}	
	return
}

func (c *Replayable_pubsub_channel) Subscribe(receiver chan string) {
	c.lazyInit()	
	c.Subscribers.PushBack(receiver)
	return 
}

func (c *Replayable_pubsub_channel) Unsubscribe(receiver chan string) {
	for e := c.Subscribers.Front(); e != nil; e = e.Next() {				
		if e.Value == receiver {
			c.Subscribers.Remove(e)
		}
	}	
	return
}