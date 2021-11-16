package main

type Queue struct {
	list []interface{}
}

func NewQueue() *Queue {
	list := make([]interface{}, 0)
	return &Queue{list}
}

func (q *Queue) Push(data interface{}) {
	q.list = append(q.list, data)
}

func (q *Queue) Pop() interface{} {
	if len(q.list) == 0 {
		return nil
	}
	head := q.list[0]
	q.list = q.list[1:]
	return head
}
