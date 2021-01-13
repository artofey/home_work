package sqlstorage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	st "github.com/artofey/home_work/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"

	// Register the DB driver.
	_ "github.com/jackc/pgx/stdlib"
)

const (
	timeFormat = "01-02-2006 15:04:05"
)

type SQLStorage struct {
	db  *sql.DB
	ctx context.Context
}

func (s *SQLStorage) AddEvent(e st.Event) (st.IDT, error) {
	if !s.permissible(e) {
		return st.IDT{}, st.ErrDateBusy
	}
	query := `
	insert into events ("user", title, descr, start_date, end_date, notify_time)
	values ($1, $2, $3, $4, $5, $6) RETURNING id
	`
	var ID string

	err := s.db.QueryRowContext(s.ctx, query, toDBConvert(e)...).Scan(&ID)
	if err != nil {
		return st.IDT{}, err
	}
	return st.IDT(uuid.MustParse(ID)), nil
}

func (s *SQLStorage) permissible(e st.Event) bool {
	dayEvents, err := s.GetDayEvents(e.DateTime)
	if err != nil {
		panic(err.Error())
	}
	for _, ev := range dayEvents {
		start := ev.DateTime
		finish := ev.DateTime.Add(ev.TimeDuration)
		// если начало или конец внутри события
		if st.TimeInsideEvent(start, e) || st.TimeInsideEvent(finish, e) {
			return false
		}
	}
	return true
}

func (s *SQLStorage) UpdateEvent(id st.IDT, e st.Event) error {
	query := `
	update events set
		"user" = $1,
		title = $2,
		descr = $3,
		start_date = $4,
		end_date = $5,
		notify_time = $6
	where id = $7
	`
	args := toDBConvert(e)
	args = append(args, id.String())
	_, err := s.db.ExecContext(s.ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update event by id = %v : %w", id.String(), err)
	}
	return nil
}

func (s *SQLStorage) DeleteEvent(id st.IDT) error {
	query := `
	delete from events where id = $1
	`
	_, err := s.db.ExecContext(s.ctx, query, id.String())
	if err != nil {
		return fmt.Errorf("failed to delete event by id = %v : %w", id.String(), err)
	}
	return nil
}

func (s *SQLStorage) DeleteAllEvents() error {
	query := `
	delete from events
	`
	_, err := s.db.ExecContext(s.ctx, query)
	if err != nil {
		return fmt.Errorf("failed to clear events: %w", err)
	}
	return nil
}

func (s *SQLStorage) GetEvent(id st.IDT) (st.Event, error) {
	query := `
	select * from events where id = $1
	`
	row := s.db.QueryRowContext(s.ctx, query, uuid.UUID(id).String())
	args := make([]interface{}, 7)
	err := row.Scan(&args[0], &args[1], &args[2], &args[3], &args[4], &args[5], &args[6])
	if errors.Is(err, sql.ErrNoRows) {
		return st.Event{}, st.ErrID
	} else if err != nil {
		return st.Event{}, err
	}
	var e st.Event = fromDBConvert(args)
	return e, nil
}

func (s *SQLStorage) getEvents(timeInterval string, dt time.Time) ([]st.Event, error) {
	query := `
	select * from events where date_trunc('%s', start_date) = date_trunc('%s', $1::timestamp)
	`
	query = fmt.Sprintf(query, timeInterval, timeInterval)
	rows, err := s.db.QueryContext(s.ctx, query, dt.Format(timeFormat))
	if err != nil {
		return []st.Event{}, err
	}
	defer rows.Close()

	var events []st.Event
	args := make([]interface{}, 9)
	for rows.Next() {
		err := rows.Scan(&args[0], &args[1], &args[2], &args[3], &args[4], &args[5], &args[6])
		if err != nil {
			return []st.Event{}, err
		}
		events = append(events, fromDBConvert(args))
	}

	if err = rows.Err(); err != nil {
		return []st.Event{}, err
	}
	return events, nil
}

func (s *SQLStorage) GetDayEvents(dt time.Time) ([]st.Event, error) {
	return s.getEvents("day", dt)
}

func (s *SQLStorage) GetWeekEvents(dt time.Time) ([]st.Event, error) {
	return s.getEvents("week", dt)
}

func (s *SQLStorage) GetMonthEvents(dt time.Time) ([]st.Event, error) {
	return s.getEvents("month", dt)
}

func (s *SQLStorage) Close() error {
	if err := s.db.Close(); err != nil {
		return fmt.Errorf("failed to close database: %w", err)
	}
	return nil
}

func New(ctx context.Context, dsn string) (st.EventsStorage, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return &SQLStorage{}, fmt.Errorf("failed to load driver: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return &SQLStorage{}, fmt.Errorf("failed to connect database: %w", err)
	}

	storageP := &SQLStorage{
		db:  db,
		ctx: ctx,
	}
	return storageP, nil
}

func toDBConvert(e st.Event) []interface{} {
	var args []interface{}

	args = append(args, uuid.UUID(e.UserID).String())   // user
	args = append(args, e.Title)                        // title
	args = append(args, e.Description)                  // descr
	args = append(args, e.DateTime)                     // start_date
	args = append(args, e.DateTime.Add(e.TimeDuration)) // end_date
	args = append(args, e.NotificationTime)             // notify_time
	return args
}

func fromDBConvert(args []interface{}) st.Event {
	startDate := args[4].(time.Time) // start_date
	endDate := args[5].(time.Time)   // end_date
	duration := endDate.Sub(startDate)
	notifyTime := args[6].(time.Time) // notify_time
	e := st.Event{
		ID:               st.IDT(uuid.MustParse(args[0].(string))), // id
		UserID:           st.IDT(uuid.MustParse(args[1].(string))), // user
		Title:            args[2].(string),                         // title
		Description:      args[3].(string),                         // descr
		DateTime:         startDate,
		TimeDuration:     duration,
		NotificationTime: notifyTime,
	}
	return e
}
