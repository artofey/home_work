package hw04_lru_cache //nolint:golint,stylecheck

// List is ...
type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	count int
	first *ListItem
	last  *ListItem
}

// Длина списка.
func (l *list) Len() int {
	return l.count
}

// Первый элемент списка.
func (l *list) Front() *ListItem {
	return l.first
}

// Последний элемент списка.
func (l *list) Back() *ListItem {
	return l.last
}

// Добавить значение в начало.
func (l *list) PushFront(v interface{}) *ListItem {
	newListItem := ListItem{Value: v}
	switch {
	case l.count == 0:
		l.first = &newListItem
		l.last = &newListItem
	case l.count == 1:
		newListItem.Next = l.first
		l.last.Prev = &newListItem
		l.first = &newListItem
	case l.count > 1:
		l.first.Prev = &newListItem
		newListItem.Next = l.first
		l.first = &newListItem
	default:
		panic("count < 0")
	}
	l.count++
	return &newListItem
}

// Добавить значение в конец.
func (l *list) PushBack(v interface{}) *ListItem {
	newListItemP := &ListItem{Value: v}
	switch {
	case l.count == 0:
		l.first = newListItemP
		l.last = newListItemP
	case l.count == 1:
		l.first.Next = newListItemP
		newListItemP.Prev = l.first
		l.last = newListItemP
	case l.count > 1:
		l.last.Next = newListItemP
		newListItemP.Prev = l.last
		l.last = newListItemP
	default:
		panic("count < 0")
	}
	l.count++
	return newListItemP
}

// Удалить элемент.
func (l *list) Remove(i *ListItem) {
	if l.count <= 0 {
		return
	}
	if l.count == 1 {
		l.first = nil
		l.last = nil
		l.count--
	} else if l.count > 1 {
		if i.Next != nil && i.Prev != nil {
			i.Prev.Next, i.Next.Prev = i.Next, i.Prev
			l.count--
		}
	}
}

// Переместить элемент в начало.
func (l *list) MoveToFront(i *ListItem) {
	if i == l.first {
		return
	}
	iPrev := i.Prev
	if i == l.last {
		// делаем предыдущий последним
		iPrev.Next = nil
		l.last = iPrev
	} else {
		// соединяем предыдущий со следующим
		iPrev.Next = i.Next
		i.Next.Prev = iPrev
	}
	// делаем элемент первым
	i.Prev = nil
	i.Next = l.first
	l.first.Prev = i
	l.first = i // перезаписали первого
}

// NewList is ...
func NewList() List {
	return &list{}
}
