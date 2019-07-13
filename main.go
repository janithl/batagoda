package main

import (
	"log"
	"math/rand"
	"os"
	"time"

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

	greetings := []string{
		"මීට කලින් මාව දැකල නැද්ද?",
		"ඇයි පකෝ?",
		"තොට ඇම්ම කියල මම පලිද?",
		"මොකද හුත්තො අනින්නෙ?",
	}

	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(greetings)
	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, greetings[n])
	})

	b.Start()
}
