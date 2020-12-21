package memorystorage

import (
	"sync"
	"time"

	st "github.com/artofey/home_work/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	// TODO
	mu sync.RWMutex
	ev st.Event
}

func (s Storage) AddEvent(e st.Event) error {
	return nil
}

func (s Storage) UpdateEvent(id st.IDT, e st.Event) error {
	return nil

}

func (s Storage) DeleteEvent(id st.IDT) error {
	return nil

}

func (s Storage) ShowDayEvents(dt time.Time) error {
	return nil

}

func (s Storage) ShowWeekEvents(dt time.Time) error {
	return nil

}

func (s Storage) ShowMonthEvents(dt time.Time) error {
	return nil

}

func New() st.EventsStorage {
	return &Storage{}
}

// TODO
