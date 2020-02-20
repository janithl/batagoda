package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
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
		appID           = os.Getenv("WOLFRAM_APPID")
		useExperimental = false
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
			if len(appID) > 0 && strings.Contains(prompt, "?") {
				var err error
				response, err = askWolfram(appID, prompt)
				if err != nil {
					response = err.Error()
				}
			} else if useExperimental {
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

func askWolfram(appID string, question string) (string, error) {
	endpoint := url.URL{
		Scheme: "https",
		Host:   "api.wolframalpha.com",
		Path:   "v1/result",
	}
	query := endpoint.Query()
	query.Set("appid", appID)
	query.Set("units", "metric")
	query.Set("i", question)
	endpoint.RawQuery = query.Encode()

	resp, err := http.Get(endpoint.String())
	if err != nil || resp.StatusCode != 200 {
		return "", errors.New("can't answer that, boss")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("can't answer that, boss")
	}

	return fmt.Sprintf("%s", body), nil
}
