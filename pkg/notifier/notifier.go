package notifier

// Notifier base interface
type Notifier interface {
	Run()
	notify(msg Message)
}

// Message is the struct for all outgoing messages
type Message struct {
	Channel       string
	Text          string
	ParseMarkdown bool
}

// NewMessage returns an Message struct for further usage
func NewMessage() Message {
	return Message{ParseMarkdown: false}
}
