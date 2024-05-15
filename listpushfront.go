package piscine

func ListPushFront(l *List, data interface{}) {
	ele := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = ele
		l.Tail = ele
		return
	}
	ele.Next = l.Head
	l.Head = ele
}
