package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	ele := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = ele
	} else {
		l.Tail.Next = ele
	}
	l.Tail = ele
}
