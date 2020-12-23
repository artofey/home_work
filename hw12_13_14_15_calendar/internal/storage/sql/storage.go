package sqlstorage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/spf13/viper"

	st "github.com/artofey/home_work/hw12_13_14_15_calendar/internal/storage"
)

var dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
	viper.GetString("db.user"),
	viper.GetString("db.password"),
	viper.GetString("db.host"),
	viper.GetString("db.port"),
	viper.GetString("db.dbname"),
)

type SQLStorage struct {
	db *sql.DB
}

func (s SQLStorage) AddEvent(e st.Event) error {
	return nil
}

func (s SQLStorage) UpdateEvent(id st.IDT, e st.Event) error {
	return nil

}

func (s SQLStorage) DeleteEvent(id st.IDT) error {
	return nil

}

func (s SQLStorage) GetDayEvents(dt time.Time) error {
	return nil

}

func (s SQLStorage) GetWeekEvents(dt time.Time) error {
	return nil

}

func (s SQLStorage) GetMonthEvents(dt time.Time) error {
	return nil

}

func (s *SQLStorage) Connect(ctx context.Context) error {
	if err := s.db.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}
	return nil
}

func (s *SQLStorage) Close(ctx context.Context) error {
	if err := s.db.Close(); err != nil {
		return fmt.Errorf("failed to close database: %w", err)
	}
	return nil
}

func New() (st.EventsStorage, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return SQLStorage{}, fmt.Errorf("failed to load driver: %w", err)
	}
	return SQLStorage{
		db: db,
	}, nil
}
