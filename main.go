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
		prompt := strings.ToLower(m.Text)
		if strings.Contains(prompt, "බටගොඩ") || strings.Contains(prompt, "batagoda") {
			prompt = strings.TrimSpace(strings.ReplaceAll(
				strings.ReplaceAll(prompt, "batagoda", ""), "බටගොඩ", ""))
			b.Send(m.Chat, liz.ReplyTo(prompt))
		}
	})

	b.Start()
}
