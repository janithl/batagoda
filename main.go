package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	liz "github.com/janithl/batagoda/eliza"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		port      = os.Getenv("PORT")
		publicURL = os.Getenv("PUBLIC_URL")
		token     = os.Getenv("TOKEN")
	)

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatalln(err)
	}

	b.Handle(tb.OnText, func(m *tb.Message) {
		rand.Seed(time.Now().Unix())
		if strings.Contains(strings.ToLower(m.Text), "බටගොඩ") ||
			strings.Contains(strings.ToLower(m.Text), "batagoda") {
			prompt := strings.ReplaceAll(strings.ToLower(m.Text), "batagoda", "")
			b.Send(m.Chat, liz.ReplyTo(prompt))
		}
	})

	b.Start()
}
