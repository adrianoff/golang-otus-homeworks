package hw04lrucache

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
	head   *ListItem
	tail   *ListItem
	length int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}
	if l.head != nil {
		newItem.Next = l.head
		l.head.Prev = newItem
	}
	if l.tail == nil {
		l.tail = newItem
	}
	l.head = newItem
	l.length++

	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}
	if l.tail != nil {
		newItem.Prev = l.tail
		l.tail.Next = newItem
	}
	if l.head == nil {
		l.head = newItem
	}
	l.tail = newItem
	l.length++

	return newItem
}

func (l *list) Remove(i *ListItem) {
	switch {
	case i.Prev != nil && i.Next != nil:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	case i.Prev == nil:
		l.head = i.Next
		l.head.Prev = nil
	case i.Next == nil:
		l.tail = i.Prev
		l.tail.Next = nil
	}

	i.Next = nil
	i.Prev = nil
	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	if i != l.head {
		l.Remove(i)
		l.PushFront(i.Value)
	}
}
