package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/artofey/home_work/hw12_13_14_15_calendar/internal/app"
	"github.com/artofey/home_work/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/artofey/home_work/hw12_13_14_15_calendar/internal/server/http"
	st "github.com/artofey/home_work/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/artofey/home_work/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/artofey/home_work/hw12_13_14_15_calendar/internal/storage/sql"
	"github.com/spf13/viper"
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

func NewEventStorage(ctx context.Context, isDB bool) (st.EventsStorage, error) {
	if isDB {
		fmt.Println("init db storage")
		DSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			viper.GetString("db.user"),
			viper.GetString("db.password"),
			viper.GetString("db.host"),
			viper.GetString("db.port"),
			viper.GetString("db.dbname"),
		)
		return sqlstorage.New(ctx, DSN)
	}
	fmt.Println("init memory storage")
	return memorystorage.New()
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	err := InitConfig()
	errHandle(err)

	logg, err := logger.New(
		viper.GetString("logger.level"),
		viper.GetString("logger.file"),
	)
	errHandle(err)
	defer func() {
		err := logg.Sync()
		errHandle(err)
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	storage, err := NewEventStorage(ctx, viper.GetBool("db.use_db"))
	errHandle(err)
	defer storage.Close()

	calendar := app.New(logg, storage)

	server := internalhttp.NewServer(calendar)

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
		errHandle(err)
	}
}
