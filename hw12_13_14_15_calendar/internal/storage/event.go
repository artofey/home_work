package storage

import (
	"fmt"
	"time"
)

var (
	ErrDateBusy = fmt.Errorf("Данное время уже занято другим событием.")
)

type IDT int64

type EventsStorage interface {
	AddEvent(e Event) error
	UpdateEvent(id IDT, e Event) error
	DeleteEvent(id IDT) error
	ShowDayEvents(dt time.Time) error
	ShowWeekEvents(dt time.Time) error
	ShowMonthEvents(dt time.Time) error
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

}
