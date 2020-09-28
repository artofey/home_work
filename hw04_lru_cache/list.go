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
	ListItem
}

// Длина списка
func (list) Len() int {
	return 0
}

// Первый элемент списка
func (list) Front() *ListItem {
	return nil
}

// Последний элемент списка
func (list) Back() *ListItem {
	return nil
}

// Добавить значение в начало
func (list) PushFront(v interface{}) *ListItem {
	return nil
}

// Добавить значение в конец
func (list) PushBack(v interface{}) *ListItem {
	return nil
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
