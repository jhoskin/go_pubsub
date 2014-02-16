package replayable_pubsub_channel

import (
	"testing"
	"time"
	"container/list"
)

func TestSubscribe(t *testing.T) {
	subscribers := list.New()

	receiver := make(chan string, 100)
	c := Replayable_pubsub_channel{ subscribers }	
	
	c.Subscribe(receiver)

	if subscribers.Len() != 1 {
		t.Fatalf("should be one subscriber, found %v", subscribers.Len())
	}

	if subscribers.Front().Value.(chan string) != receiver {
		t.Fatal("receiver should be subscribed")
	}
}

func TestPublish(t *testing.T) {
	c := Replayable_pubsub_channel{}
	c.Publish("test")
}

func TestEndToEnd(t *testing.T) {
	receiver := make(chan string, 100)
	c := Replayable_pubsub_channel{}
	
	c.Subscribe(receiver)

	var testMessage string = "test message"
	c.Publish(testMessage)

	timeout := make(chan bool, 1)
	go func() {
	    time.Sleep(1 * time.Second)
	    timeout <- true
	}()

	var result string
	select {
		case result = <-receiver:			
			if result != testMessage {
				t.Errorf("wrong message received %v", result)				
			}
		case <-timeout:
			t.Errorf("timed out waiting for message")
	}	

}

func TestUnsubscribe(t *testing.T) {
	subscribers := list.New()

	receiver := make(chan string, 100)
	c := Replayable_pubsub_channel{ subscribers }	
	
	c.Subscribe(receiver)
	
	if subscribers.Len() != 1 {
		t.Fatalf("should be one subscriber, found %v", subscribers.Len())
	}

	if subscribers.Front().Value.(chan string) != receiver {
		t.Fatal("receiver should be subscribed")
	}

	c.Unsubscribe(receiver)		

	if subscribers.Len() != 0 {
		t.Fatalf("should be no subscribers, found %v", subscribers.Len())
	}
}