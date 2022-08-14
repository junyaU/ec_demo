package adapter

type MessagePublisher interface {
	Publish(topic string, args ...interface{}) error
}
