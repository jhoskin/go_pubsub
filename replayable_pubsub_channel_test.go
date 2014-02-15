package replayable_pubsub_channel

import (
	"testing"
	"time"
	"container/list"
)

func TestSubscribe(t *testing.T) {
	subscribers := list.New()

	receiver := make(chan string, 100)
	c := Replayable_pubsub_channel{subscribers}
	c.Subscribe(receiver)

	err := c.Subscribe(receiver)
	if err != nil {
		t.Fatal(err)
	}

	if subscribers.Len() != 1 {
		t.Fatal("should be one subscriber")
	}

	if subscribers.Front().(chan string) != receiver {
		t.Fatal("receiver should be subscribed")
	}
}

func TestPublish(t *testing.T) {
	c := Replayable_pubsub_channel{}
	err := c.Publish("test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestEndToEnd(t *testing.T) {
	subscribers := list.New()

	receiver := make(chan string, 100)
	c := Replayable_pubsub_channel{subscribers}
	c.Subscribe(receiver)

	err := c.Subscribe(receiver)
	if err != nil {
		t.Error(err)
	}

	err = c.Publish("test")
	if err != nil {
		t.Error(err)
	}

	timeout := make(chan bool, 1)
	go func() {
	    time.Sleep(1 * time.Second)
	    timeout <- true
	}()

	var result string
	select {
		case result = <-receiver:
			if result != "test" {
				t.Errorf("wrong message received %v", result)
			}
		case <-timeout:
			t.Errorf("timed out waiting for message")
	}	
}
