package telegram

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mmatveyev/homebot/common"
	"log"
	"math"
	"regexp"
	"time"
)

var wash1 = regexp.MustCompile("(?i)(прання)[^\\d]*((\\d+).?(\\d*))*")
var loc *time.Location

type tb struct {
	bot *tgbotapi.BotAPI
}

func NewBot(cfg common.Telegram) error {
	thebot := tb{}
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = cfg.Debug
	thebot.bot = bot
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	loc, err = time.LoadLocation(cfg.ClientTimezone)
	if err != nil {
		return err
	}
	updates, err := thebot.bot.GetUpdatesChan(u)

	go thebot.serve(updates)
	return err
}
func (thebot *tb) serve(updates tgbotapi.UpdatesChannel) {
	var msg tgbotapi.MessageConfig

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if matches := wash1.FindStringSubmatch(update.Message.Text); len(matches) > 0 {
			now := time.Now()
			yy, mm, dd := now.Date()
			today := time.Date(yy, mm, dd, 0, 0, 0, 0, loc)

			morning := today.Add(time.Hour * 7)
			if now.After(morning) {
				morning = morning.Add(time.Hour * 24)
			}
			morningDelay := math.Trunc(morning.Sub(now).Hours())
			text := fmt.Sprintf("Став затримку %.f годин\n", morningDelay)

			if matches[3] != "" {
				if matches[4] == "" {
					matches[4] = "0"
				}
				dur, err := time.ParseDuration(matches[3] + "h" + matches[4] + "m")
				if err == nil {
					todayEvening := today.Add(time.Hour * 23)
					eveningDelay := math.Ceil(todayEvening.Add(dur).Sub(now).Hours())
					text += fmt.Sprintf("Або %.f Годин якщо ввечері", eveningDelay)
				}
			}

			msg = tgbotapi.NewMessage(update.Message.Chat.ID, text)
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		}

		_, err := thebot.bot.Send(msg)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
