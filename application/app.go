package application

import (
	"flag"
	"github.com/mmatveyev/homebot/common"
	"github.com/mmatveyev/homebot/telegram"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var signalChan = make(chan os.Signal, 1)

func Run() int {
	var configPath string
	flag.StringVar(&configPath, "c", "config.toml", "Path to configuration file")
	flag.Parse()

	config, err := common.NewConfigFromFile(configPath)
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
