package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	batz "github.com/janithl/batagoda/batagodax"
	liz "github.com/janithl/batagoda/eliza"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		port            = os.Getenv("PORT")
		publicURL       = os.Getenv("PUBLIC_URL")
		token           = os.Getenv("TOKEN")
		useExperimental = true
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
			response := ""
			if useExperimental {
				response = batz.Respond(prompt)
			} else {
				response = liz.ReplyTo(prompt)
			}

			b.Send(m.Chat, response)
		}
	})

	b.Handle("/switch", func(m *tb.Message) {
		useExperimental = !useExperimental
		if useExperimental {
			b.Send(m.Sender, "batagodax mode selected")
		} else {
			b.Send(m.Sender, "eliza mode selected")
		}

	})

	b.Start()
}
