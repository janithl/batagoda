package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	liz "github.com/kennysong/goeliza"
	tb "gopkg.in/tucnak/telebot.v2"
)

var greetings = []string{
	"මීට කලින් මාව දැකල නැද්ද?",
	"ඇයි පකෝ?",
	"තොට ඇම්ම කියල මම පලිද?",
	"මොකද හුත්තො අනින්නෙ?",
	"තොගෙ අඳෝනාව අහන් ඉඳල මට බඩ යනව",
	"මේකට වඩා හොඳයි ඇට වල මවිල් ගනං කරන එක",
	"ගිහිං වැලක් බලපංකො පව් නොදී",
	"තොපිගෙ මුල් ලමාවිය ප්‍රශ්න තමයි පෙන්නන්නෙ",
}

func greet() string {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(len(greetings))
	return greetings[n]
}

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
		if strings.Contains(strings.ToLower(m.Text), "batagoda") ||
			strings.Contains(strings.ToLower(m.Text), "බටගොඩ") {
			str := liz.ReplyTo(m.Text)
			b.Send(m.Chat, str)
		}
	})

	b.Start()
}
