package storage

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrDateBusy = fmt.Errorf("данное время уже занято другим событием")
	ErrID       = fmt.Errorf("ошибочный идентификатор события")
)

// IDT is UUID type.
type IDT uuid.UUID

func (i *IDT) String() string {
	return uuid.UUID(*i).String()
}

type EventsStorage interface {
	AddEvent(e Event) (IDT, error)
	GetEvent(id IDT) (Event, error)
	UpdateEvent(id IDT, e Event) error
	DeleteEvent(id IDT) error
	DeleteAllEvents() error
	GetDayEvents(dt time.Time) ([]Event, error)
	GetWeekEvents(dt time.Time) ([]Event, error)
	GetMonthEvents(dt time.Time) ([]Event, error)
	Close() error
}

type Event struct {
	ID               IDT
	UserID           IDT
	Title            string
	Description      string
	DateTime         time.Time
	TimeDuration     time.Duration
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
		Title:            title,
		DateTime:         dateTime,
		TimeDuration:     timeDuration,
		UserID:           userID,
		NotificationTime: dateTime.Add(-1 * time.Hour),
	}
}

func TimeInsideEvent(t time.Time, e Event) bool {
	// если время после начала и до конца
	if t.After(e.DateTime) && t.Before(e.DateTime.Add(e.TimeDuration)) {
		return true
	}
	return false
}
