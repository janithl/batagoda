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
		"තොගෙ අඳෝනාව අහන් ඉඳල මට බඩ යනව",
		"මේකට වඩා හොඳයි ඇට වල මවිල් ගනං කරන එක",
		"ගිහිං වැලක් බලපංකො පව් නොදී",
		"තොපිගෙ මුල් ලමාවිය ප්‍රශ්න තමයි පෙන්නන්නෙ",
	}

	b.Handle("@BatagodaBot", func(m *tb.Message) {
		rand.Seed(time.Now().Unix())
		n := rand.Intn(len(greetings))
		b.Send(m.Chat, greetings[n])
	})

	b.Start()
}
