package main

import (
	"log"
	"time"

	"github.com/subosito/gotenv"
	tele "gopkg.in/telebot.v3"
)

func init() {
	_ = gotenv.Load()
}

func main() {
	pref := tele.Settings{
		Token:  "5976194824:AAGfIdPtCeo04j3JsOudLy6DlsQjC9u2n5A",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Start()
}
