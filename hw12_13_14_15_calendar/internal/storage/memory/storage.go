package memorystorage

import (
	"sync"
	"time"

	st "github.com/artofey/home_work/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
)

type MemoryStorage struct {
	mu sync.RWMutex
	st map[st.IDT]st.Event
}

// AddEvent - добавить новое событие в хранилище.
func (s *MemoryStorage) AddEvent(e st.Event) (st.IDT, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.permissible(e) {
		return st.IDT{}, st.ErrDateBusy
	}
	var emptyID st.IDT
	if e.ID != emptyID {
		s.st[e.ID] = e
		return e.ID, nil
	}
	// если идентификатора нет, то генерим новый и записываем в событие.
	var newID st.IDT
	for {
		newID = st.IDT(uuid.New())
		if _, ok := s.st[newID]; !ok {
			break
		}
	}
	e.ID = newID

	s.st[newID] = e
	return newID, nil
}

// GetEvent - получить событие по идентификатору.
func (s *MemoryStorage) GetEvent(id st.IDT) (st.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	event, ok := s.st[id]
	if !ok {
		return st.Event{}, st.IDError
	}
	return event, nil
}

// UpdateEvent - обновить событие в хранилище.
func (s *MemoryStorage) UpdateEvent(id st.IDT, e st.Event) error {
	event, err := s.GetEvent(id)
	if err != nil {
		return err
	}
	event.DateTime = e.DateTime
	event.Description = e.Description
	event.NotificationTime = e.NotificationTime
	event.TimeDuration = e.TimeDuration
	event.Title = e.Title
	event.UserID = e.UserID
	_, err = s.AddEvent(event)
	if err != nil {
		return err
	}
	return nil
}

// DeleteEvent - удалить событие из хранилища.
func (s *MemoryStorage) DeleteEvent(id st.IDT) error {
	_, err := s.GetEvent(id)
	if err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.st, id)
	return nil
}

// GetDayEvents - получить список событий за день.
func (s *MemoryStorage) GetDayEvents(dt time.Time) ([]st.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	y := dt.Year()
	d := dt.YearDay()
	dayEvents := make([]st.Event, 0)
	for _, ev := range s.st {
		if ev.DateTime.Year() == y && ev.DateTime.YearDay() == d {
			dayEvents = append(dayEvents, ev)
		}
	}
	return dayEvents, nil
}

// GetWeekEvents - получить список событий за неделю.
func (s *MemoryStorage) GetWeekEvents(dt time.Time) ([]st.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	y, w := dt.ISOWeek()
	weekEvents := make([]st.Event, 0)
	for _, ev := range s.st {
		year, week := ev.DateTime.ISOWeek()
		if year == y && week == w {
			weekEvents = append(weekEvents, ev)
		}
	}
	return weekEvents, nil
}

// GetMonthEvents - получить список событий за месяц.
func (s *MemoryStorage) GetMonthEvents(dt time.Time) ([]st.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	y := dt.Year()
	m := dt.Month()
	monthEvents := make([]st.Event, 0)
	for _, ev := range s.st {
		if ev.DateTime.Year() == y && ev.DateTime.Month() == m {
			monthEvents = append(monthEvents, ev)
		}
	}
	return monthEvents, nil
}

// Проверка не пересекается ли событие с уже имеющимися.
func (s *MemoryStorage) permissible(e st.Event) bool {
	for _, ev := range s.st {
		start := ev.DateTime
		finish := ev.DateTime.Add(ev.TimeDuration)
		// если начало или конец внутри события
		if timeInsideEvent(start, e) || timeInsideEvent(finish, e) {
			return false
		}
	}
	return true
}

func timeInsideEvent(t time.Time, e st.Event) bool {
	// если время после начала и до конца
	if t.After(e.DateTime) && t.Before(e.DateTime.Add(e.TimeDuration)) {
		return true
	}
	return false
}

// New - создает новый инстанс хранилища.
func New() (st.EventsStorage, error) {
	return &MemoryStorage{
		st: make(map[st.IDT]st.Event),
	}, nil
}
