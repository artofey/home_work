package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/viper"

	"github.com/artofey/home_work/hw12_13_14_15_calendar/internal/app"
	"github.com/artofey/home_work/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/artofey/home_work/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/artofey/home_work/hw12_13_14_15_calendar/internal/storage/memory"
)

const defaultConfigFile = "/etc/calendar/config.toml"

var configFile string

func init() {
	flag.StringVar(&configFile, "config", defaultConfigFile, "Path to configuration file")
}

func errHandle(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	err := InitConfig()
	errHandle(err)

	logg, err := logger.New(viper.GetString("logger.level"))
	errHandle(err)
	defer logg.Sync()

	storage := memorystorage.New()
	calendar := app.New(logg, storage)

	server := internalhttp.NewServer(calendar)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals)

		<-signals
		signal.Stop(signals)
		cancel()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		os.Exit(1)
	}
}
