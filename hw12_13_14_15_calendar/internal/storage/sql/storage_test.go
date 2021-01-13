package sqlstorage

import (
	"context"
	"errors"
	"testing"
	"time"

	st "github.com/artofey/home_work/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestSQLStorage(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s, err := New(ctx, "postgres://calendar_app:12345678@localhost:54321/calendar")
	require.NoError(t, err)
	err = s.DeleteAllEvents()
	require.NoError(t, err)

	d, _ := time.ParseDuration("2h")
	userID := st.IDT(uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))

	ev1_12 := st.NewEvent(
		"Тестовое событие 1 февраля 12:00",
		time.Date(2021, 2, 1, 12, 0, 0, 0, time.UTC),
		d,
		userID,
	)
	// потытка получить событие по невалидному идентификатору
	_, err = s.GetEvent(userID)
	require.True(t, errors.Is(err, st.ErrID))

	id, err := s.AddEvent(ev1_12)
	require.NoError(t, err)

	// поулчили только что добавленное событие
	event, err := s.GetEvent(id)
	require.Equal(t, ev1_12.Title, event.Title)

	ev1_13 := st.NewEvent(
		"Тестовое событие 1 февраля 13:00",
		time.Date(2021, 2, 1, 13, 0, 0, 0, time.UTC),
		d,
		userID,
	)

	// при попытке добавить событие с временем пересекающимся с имеющимися событиями возникает ошибка
	_, err = s.AddEvent(ev1_13)
	require.Equal(t, st.ErrDateBusy, err)

	ev1_11 := st.NewEvent(
		"Тестовое событие 1 февраля 11:00",
		time.Date(2021, 2, 1, 11, 0, 0, 0, time.UTC),
		d,
		userID,
	)

	// то же самое но для события пересекающегося с началом имеющегося
	_, err = s.AddEvent(ev1_11)
	require.Equal(t, err, st.ErrDateBusy)

	ev1_15 := st.NewEvent(
		"Тестовое событие 1 февраля 15:00",
		time.Date(2021, 2, 1, 17, 0, 0, 0, time.UTC),
		d,
		userID,
	)
	_, err = s.AddEvent(ev1_15)
	require.NoError(t, err)
	events, err := s.GetDayEvents(time.Date(2021, 2, 1, 15, 0, 0, 0, time.UTC))
	require.NoError(t, err)
	// после добавления еще одного события получаем два события в этот день.
	require.Equal(t, 2, len(events))

	ev7_00 := st.NewEvent(
		"Тестовое событие 7 февраля 00:00",
		time.Date(2021, 2, 7, 0, 0, 0, 0, time.UTC),
		d,
		userID,
	)
	_, err = s.AddEvent(ev7_00)
	// в конец недели добавляем еще одно событие.
	require.NoError(t, err)
	weekEvents, err := s.GetWeekEvents(time.Date(2021, 2, 1, 15, 0, 0, 0, time.UTC))
	require.NoError(t, err)
	// на неделе теперь должно быть три события
	require.Equal(t, 3, len(weekEvents))
	dayEvents, err := s.GetDayEvents(time.Date(2021, 2, 1, 15, 0, 0, 0, time.UTC))
	require.NoError(t, err)
	// но в тот же день только два
	require.Equal(t, 2, len(dayEvents))

	ev16_00 := st.NewEvent(
		"Тестовое событие 16 февраля 00:00",
		time.Date(2021, 2, 16, 0, 0, 0, 0, time.UTC),
		d,
		userID,
	)
	id16_00, err := s.AddEvent(ev16_00)
	// в середине месяца добавляем еще одно событие.
	require.NoError(t, err)
	monthEvents, err := s.GetMonthEvents(time.Date(2021, 2, 1, 15, 0, 0, 0, time.UTC))
	require.NoError(t, err)
	// итого всего за месяц должно быть 4 события
	require.Equal(t, 4, len(monthEvents))

	ev20_00 := st.NewEvent(
		"Тестовое событие 20 февраля 00:00",
		time.Date(2021, 2, 20, 0, 0, 0, 0, time.UTC),
		d,
		userID,
	)
	// добивили еще одно событие
	id, err = s.AddEvent(ev20_00)
	require.NoError(t, err)
	monthEvents, err = s.GetMonthEvents(time.Date(2021, 2, 1, 15, 0, 0, 0, time.UTC))
	require.NoError(t, err)
	// теперь их в месяце 5 штук
	require.Equal(t, 5, len(monthEvents))
	err = s.DeleteEvent(id)
	require.NoError(t, err)
	monthEvents, err = s.GetMonthEvents(time.Date(2021, 2, 1, 15, 0, 0, 0, time.UTC))
	// после удаления опять 4 события
	require.Equal(t, 4, len(monthEvents))

	// создать новое событие
	userID2 := st.IDT(uuid.MustParse("7ba7b810-9dad-11d1-80b4-00c04fd430c8"))
	d, _ = time.ParseDuration("5h")
	ev16_14 := st.NewEvent(
		"Тестовое обновленное событие 16 февраля 14:00",
		time.Date(2021, 2, 16, 14, 0, 0, 0, time.UTC),
		d,
		userID2,
	)
	// обновить полученное событие данными из вновь созданного
	err = s.UpdateEvent(id16_00, ev16_14)
	require.NoError(t, err)
	day16Events, err := s.GetDayEvents(time.Date(2021, 2, 16, 0, 0, 0, 0, time.UTC))
	require.NoError(t, err)
	// после обновления событие также осталось одно в этот день
	require.Equal(t, 1, len(day16Events))
	event16 := day16Events[0]
	// поля события корректно обновились
	require.Equal(t, ev16_14.Title, event16.Title)
	require.Equal(t, ev16_14.DateTime, event16.DateTime)
	require.Equal(t, ev16_14.TimeDuration, event16.TimeDuration)
	require.Equal(t, ev16_14.UserID, event16.UserID)
}
