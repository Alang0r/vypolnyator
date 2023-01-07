package telegram

import (
	"fmt"
	"regexp"
	"time"

	"github.com/Alang0r/vypolnyator/pkg/service"
	tele "gopkg.in/telebot.v3"
)

var handlers map[string]service.Request

func init() {
	handlers = make(map[string]service.Request)
}

type TeleHandler struct {
	Func tele.HandlerFunc
	MW   tele.MiddlewareFunc
}

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

func RegisterHandler(hName string, hFunc service.Request) {
	handlers[hName] = hFunc
}

func (c *Communicator) Listen() {

	c.Handle(tele.OnText, func(c tele.Context) error {
		if _, ok := handlers[c.Text()]; ok {
			code, data := service.SendRequest(handlers[c.Text()], "http://localhost:3001")
			prepRpl := fmt.Sprintf("%s returned %s  with code %s",c.Text(), data, code)
			return c.Send(prepRpl)
		}
		return c.Send("Ne-a!")
	})

	c.Start()
}

func (c *Communicator) Send() {

}

func verifyRequest(message string) () {

	// Check if message is request
	match, err := regexp.MatchString(handlerRegexp, message)
	if err != nil{
		
	}

	if match {
		// if h, ok := handlers[message]; ok {

		// 	// h.Execute()
		// }
	}

	
}