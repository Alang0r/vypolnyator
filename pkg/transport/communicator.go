package transport

type Communicator interface {
	SendMessage(mes Message) error
	Listen()
}

type Message struct {
	Body string
}

