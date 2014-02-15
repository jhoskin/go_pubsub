package replayable_pubsub_channel

import "errors"

type Replayable_pubsub_channel struct
{
	subscribers []chan string
}

func (c Replayable_pubsub_channel) Init() {
	c.subscribers = make([]chan string, 0, 10)
}

func (c Replayable_pubsub_channel) Publish(message string) (err error) {
	for _,receiver := range c.subscribers {
		receiver <- message
	}
	return
}

func (c Replayable_pubsub_channel) Subscribe(receiver chan string) (err error) {
	c.subscribers = append(c.subscribers, receiver)
	return
}

func (c Replayable_pubsub_channel) Unsubscribe(receiver chan string) (err error) {
	return errors.New("Not implemented")
}