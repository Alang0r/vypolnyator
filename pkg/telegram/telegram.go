package telegram

import (
	"fmt"
	"regexp"
	"time"

	err "github.com/Alang0r/vypolnyator/pkg/error"
	tele "gopkg.in/telebot.v3"
)

var handlers map[string]TeleHandler

func init() {
	handlers = make(map[string]TeleHandler)
}

type TeleHandler interface {
	Execute() (string, err.Error)
}

type TeleHandlerResult interface {

}

type Communicator struct {
	*tele.Bot
}

func NewCommunicator(params map[string]string) (*Communicator, err.Error) {
	c := Communicator{}

	pref := tele.Settings{
		Token:  params[ParamTgToken],
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	c.Bot, _ = tele.NewBot(pref)

	return &c, *err.New().SetCode(0)
}

func RegisterHandler(hName string, hFunc TeleHandler) {
	handlers[hName] = hFunc
}

func (c *Communicator) Listen() {

	c.Handle(tele.OnText, func(c tele.Context) error {
		if h, ok := handlers[c.Text()]; ok {
			rsp, err := h.Execute()
			if err.Code != 0 {

			}

			//return c.Send(rsp)
			
			prepRpl := fmt.Sprintf("%s returned %s  with code %d",c.Text(), rsp, err.Code)
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