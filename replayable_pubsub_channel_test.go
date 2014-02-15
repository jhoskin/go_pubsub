package replayable_pubsub_channel

import "testing"

func TestSubscribe(t *testing.T) {
	receiver := make(chan string, 100)
	c := Replayable_pubsub_channel{}
	c.Subscribe(receiver)

	err := c.Subscribe(receiver)
	if err != nil {
		t.Error(err)
	}

	err = c.Publish("test")
	if err != nil {
		t.Error(err)
	}

	result := <-receiver
	if result != "test" {
		t.Errorf("wrong message received %v", result)
	}
}


func TestPublish(t *testing.T) {
	c := Replayable_pubsub_channel{}
	err := c.Publish("test")
	if err != nil {
		t.Error(err)
	}
}
