package tools

type MessageQueue struct {
	stop chan struct{}
	errs chan error
	data []string
}

func NewMessageQueue() *MessageQueue {
	mq := MessageQueue{
		stop: make(chan struct{}, 1),
		errs: make(chan error, 1),
		data: make([]string, 1),
	}
}

func (q *MessageQueue) EnQueue(mes string) {
	q.data = append(q.data, mes)
}

func (q *MessageQueue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *MessageQueue) DeQueue() string {
	mes := q.data[0]
	q.data = q.data[1:]
	return mes
}
