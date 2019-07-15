# Batagoda

A profanity-laden Telegram bot with two distinct engines.

Uses the amazing [Telebot](https://github.com/tucnak/telebot) framework for the bot integration stuff.

## Setting Up

- Clone repo
- Copy `.env.example` to `.env` and set the values (get Telegram token from
  [BotFather](https://telegram.me/BotFather))
- `go run .` 🤷🏽‍♀️

## Why Batagoda?

The name is a throwback to the original Slack bot _Batagoda_ (may he RIP), written by Viren Dias and
Kaveen Rodrigo at [YAMU](https://www.yamu.lk) circa 2016/17.

## Engines

- Eliza: Code copied wholesale from [goeliza](https://github.com/kennysong/goeliza)
  and translated to Sinhala via Google Translate, with none of the change in sentence
  structures etc. to make it sound less awkward in Sinhala.

- BatagodaX: Written from scratch. Less sophisticated than Eliza but seems to work okay.
  Has better tests!

## License

Kenny Song's original code does not have a license, so I don't know how appropriate this is, but I'm
releasing my modifications and additions under the MIT license.
