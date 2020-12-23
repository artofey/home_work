package storage

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrDateBusy = fmt.Errorf("Данное время уже занято другим событием.")
	IDError     = fmt.Errorf("Ошибочный идентификатор события.")
)

type IDT uuid.UUID

type EventsStorage interface {
	AddEvent(e Event) (IDT, error)
	GetEvent(id IDT) (Event, error)
	UpdateEvent(id IDT, e Event) error
	DeleteEvent(id IDT) error
	GetDayEvents(dt time.Time) ([]Event, error)
	GetWeekEvents(dt time.Time) ([]Event, error)
	GetMonthEvents(dt time.Time) ([]Event, error)
}

type Event struct {
	ID               IDT
	Title            string
	DateTime         time.Time
	TimeDuration     time.Duration
	Description      string
	UserID           IDT
	NotificationTime time.Time
}

func (e *Event) AddDescription(d string) {
	e.Description = d
}

func (e *Event) SetNotificationTime(t time.Time) {
	e.NotificationTime = t
}

func NewEvent(
	title string,
	dateTime time.Time,
	timeDuration time.Duration,
	userID IDT,
) Event {
	return Event{
		Title:        title,
		DateTime:     dateTime,
		TimeDuration: timeDuration,
		UserID:       userID,
	}
}
