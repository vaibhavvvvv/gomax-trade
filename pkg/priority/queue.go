package priority

import (
	"errors"
	"sort"
)

// a priority queue for messages
type Queue struct {
	messages []*Message
}

// to create a new priority queue
func  NewQueue() *Queue {
	return &Queue{
		messages: make([]*Message,0),
	}
}

//to add a message to the queue
func (q *Queue) Push(msg *Message) error{
	if msg == nil{
		return errors.New("cannot push because message is nil")
	}
	q.messages = append(q.messages, msg)

	sort.Slice(q.messages, func(i,j int) bool{
		if  q.messages[i].Priority == q.messages[j].Priority {
			//if priority is same then the message the arrived earlier i.e. has shorted timestamp should be ahead in queue
			return q.messages[i].Timestamp < q.messages[j].Timestamp
		}
		return q.messages[i].Priority > q.messages[j].Priority
	})

	return nil
}

// to remove and return the highest priority message
func (q *Queue) Pop() (*Message, error){
	if len(q.messages) == 0 {
		return nil, errors.New("queue is empty")
	}	

	msg := q.messages[0]

	q.messages = q.messages[1:]

	return msg, nil
}

// to get the number of messages in the queue
func (q *Queue) Size() int{
	return len(q.messages)
}

// to check if queue is empty
func (q *Queue) IsEmpty() bool{
	return len(q.messages) == 0
}