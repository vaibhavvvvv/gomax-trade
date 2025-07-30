package priority

//priority levels for messages
type Priority int

const (
	Low Priority = iota
	Medium
	High
)

//message :- trading msgs with priority
type Message struct { 
	ID string
	Topic string
	Priority Priority
	Timestamp int64
	Payload []byte
}