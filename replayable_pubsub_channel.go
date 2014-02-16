package replayable_pubsub_channel

import (
	"errors"
	"container/list"
	"log"
	//"strconv"
)

type Replayable_pubsub_channel struct
{
	Subscribers list.List
}

func (c Replayable_pubsub_channel) lazyInit() {
	log.Printf("LazyInit")
	if &c.Subscribers == nil {
		log.Printf("creating new subscribers list")
		c.Subscribers = *list.New()
	}	
}

func (c Replayable_pubsub_channel) Publish(message string) (err error) {
	log.Print("Called Publish")
	c.lazyInit()
	log.Printf("number of subscribers to publish to = %v", c.Subscribers.Len())
	for e := c.Subscribers.Front(); e != nil; e.Next() {
		log.Printf("sending message")
		e.Value.(chan string) <- message
	}	
	return
}

func (c Replayable_pubsub_channel) Subscribe(receiver chan string, subs list.List) (s list.List, err error) {
	log.Print("Called Subscribe")
	log.Printf("lists are equal 1 = %v", &subs == &c.Subscribers)
	log.Printf("lists are equal 1 = %v", &subs == &c.Subscribers)
	c.lazyInit()
	log.Printf("lists are equal 2 = %v", &subs == &c.Subscribers)
	c.Subscribers.PushBack(receiver)
	log.Printf("lists are equal 3 = %v", &subs == &c.Subscribers)
	log.Printf("number of subscribers = %v", c.Subscribers.Len())
	return c.Subscribers, nil
}

func (c Replayable_pubsub_channel) Unsubscribe(receiver chan string) (err error) {
	return errors.New("Not implemented")
}