package priority

import (
	"testing"
	"time"
)

func TestNewQueue(t *testing.T) {
	q :=  NewQueue()

	if !q.IsEmpty(){
		t.Error("new queue should be empty")
	}
}

func TestPushAndPop(t *testing.T){
	q := NewQueue()

	msg:= &Message{
		ID:       "test-1",
        Topic:    "test-topic",
        Priority: Medium,
        Payload:  []byte("some test data"),
        Timestamp: time.Now().Unix(),
	}

	err := q.Push(msg)
	if err != nil {
		t.Errorf("Push operation failed: %v", err)
	}

	if q.IsEmpty() {
		t.Error("Queue shouldn't be empty after Push operation")
	}

	poppedMsg, err := q.Pop() 
	if err != nil {
		t.Errorf("Pop operation failed : %v", err)
	}

	if poppedMsg.ID != msg.ID {
		t.Errorf("Expected message ID %s but got %s", msg.ID, poppedMsg.ID)
	}

	if !q.IsEmpty() {
		t.Error("Queue should be empty after Push operation")
	}

	q.Push(&Message{ID: "low", Priority: Low, Timestamp: 1})
	q.Push(&Message{ID: "high", Priority: High, Timestamp: 2})
	q.Push(&Message{ID: "medium", Priority: Medium, Timestamp: 3})
	q.Push(&Message{ID: "high", Priority: High, Timestamp: 4})

	poppedMsg, err = q.Pop()
	if poppedMsg.ID != "high"{
		t.Errorf("Expected ID=high but got ID = %s", poppedMsg.ID)
	}
	if poppedMsg.Timestamp != 2{
		t.Errorf("Expected Timestamp=2 but got Timestamp = %v", poppedMsg.Timestamp)
	}

}