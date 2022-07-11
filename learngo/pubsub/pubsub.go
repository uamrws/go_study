package pubsub

type (
	subscriber chan interface{}
	topicFunc  func(v interface{}) bool
)
