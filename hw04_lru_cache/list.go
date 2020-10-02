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

// Длина списка
func (l *list) Len() int {
	return l.count
}

// Первый элемент списка
func (l *list) Front() *ListItem {
	return l.first
}

// Последний элемент списка
func (l *list) Back() *ListItem {
	return l.last
}

// Добавить значение в начало
func (l *list) PushFront(v interface{}) *ListItem {
	newListItem := ListItem{Value: v}
	if l.count == 0 {
		l.first = &newListItem
		l.last = &newListItem
		l.count++
		return &newListItem
	}
	l.first.Prev = &newListItem

	newListItem.Next = l.first
	l.first = &newListItem
	l.count++
	return &newListItem
}

// Добавить значение в конец
func (l *list) PushBack(v interface{}) *ListItem {
	newListItem := ListItem{Value: v}
	if l.count == 0 {
		l.first = &newListItem
		l.last = &newListItem
		l.count++
		return &newListItem
	}
	l.last.Next = &newListItem

	newListItem.Prev = l.last
	l.last = &newListItem
	l.count++
	return &newListItem
}

// Удалить элемент
func (list) Remove(i *ListItem) {

}

// Переместить элемент в начало
func (list) MoveToFront(i *ListItem) {

}

// NewList is ...
func NewList() List {
	return &list{}
}
