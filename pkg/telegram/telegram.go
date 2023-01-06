package telegram

import (
	"time"

	tele "gopkg.in/telebot.v3"
)

type Communicator struct {
	*tele.Bot
}

func NewCommunicator(params map[string]string) (*Communicator, error) {
	c := Communicator{}

	pref := tele.Settings{
		Token:  params[ParamTgToken],
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	c.Bot, _ = tele.NewBot(pref)

	return &c, nil
}

func (c *Communicator) Listen() {
	c.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})
	c.Start()
}

func (c *Communicator) Send() {

}
