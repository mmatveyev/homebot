package application

import (
	"github.com/mmatveyev/homebot/common"
	"github.com/mmatveyev/homebot/telegram"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var signalChan = make(chan os.Signal, 1)

func Run() int {
	config, err := common.GetConfig()
	if err != nil {
		log.Fatal(err.Error())
		return 1
	}
	err = telegram.NewBot(config.Telegram)
	if err != nil {
		log.Fatal(err.Error())
		return 1
	}

	HandleStopEvent()
	return 0
}

func HandleStopEvent() {
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//Wait for system signal
	_ = <-signalChan
}
