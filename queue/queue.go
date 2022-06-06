package queue

import "log"

type Queue struct {
	data []any
}

var logger = log.Default()

func (queue *Queue) Enqueue(item any) {
	queue.data = append(queue.data, item)
}

func (queue *Queue) Dequeue() (*any, bool) {
	if len(queue.data) == 0 {
		logger.Print("Queue is empty")
		return nil, false
	}
	item := queue.data[0]
	queue.data = queue.data[1:len(queue.data)]
	return &item, true
}

func (queue *Queue) Front() *any {
	item := queue.data[0]
	return &item
}

func (queue *Queue) IsEmpty() bool {
	return len(queue.data) == 0
}

func (queue *Queue) Size() int {
	return len(queue.data)
}
