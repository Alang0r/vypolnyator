package telegram

import (
	"fmt"
	"regexp"
	"time"

	err "github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/service"
	tele "gopkg.in/telebot.v3"
)

var THandlers map[string]THandler

func init() {
	THandlers = make(map[string]THandler)
}

type THandler interface {
	Run() (string, err.Error)
}

type TResult interface {
}

type Bot struct {
	tH map[string]THandler
	*tele.Bot
}

func NewBot(s *service.Service) (*Bot, err.Error) {
	b := Bot{}

	t := s.GetEnvVariable(ParamTgToken)
	pref := tele.Settings{
		Token:  t,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b.Bot, _ = tele.NewBot(pref)
	b.tH = make(map[string]THandler)
	b.tH = THandlers

	return &b, *err.New().SetCode(0)
}

func RegisterHandler(hName string, hFunc THandler) {
	THandlers[hName] = hFunc
}

func (b *Bot) Listen() {

	b.Handle(tele.OnText, func(c tele.Context) error {
		rText := c.Text()

		if h, ok := b.tH[rText]; ok {
			rsp, err := h.Run()
			if err.Code != 0 {
				return c.Send(fmt.Sprintf("error run request: %d, %s", err.Code, err.Description))
			}

			prepRpl := fmt.Sprintf("%s returned %s  with code %d", c.Text(), rsp, err.Code)
			return c.Send(prepRpl)
		} else {
			return c.Send("Not Found")
		}

		
	})

	b.Start()
}

func (c *Bot) Send() {

}

func verifyRequest(message string) {

	// Check if message is request
	match, err := regexp.MatchString(handlerRegexp, message)
	if err != nil {

	}

	if match {
		if h, ok := THandlers[message]; ok {
			h.Run()
		}
	}

}
