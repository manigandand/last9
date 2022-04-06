package task

import "last9/task/worker"

type Worker interface {
	Do()
}

type Task struct {
	event EventType
}

type EventType string

const (
	EventTypeNewInstance EventType = "new_instance"
	EventTypeNewVPC      EventType = "new_vpc"
	EventTypeNewSubnet   EventType = "new_subnet"
)

var workers = make(map[EventType][]Worker)

func Init() {
	// register workers
	workers[EventTypeNewInstance] = []Worker{
		worker.NewSSH(),
	}
}

func New(event EventType) *Task {
	return &Task{
		event: event,
	}
}

func (t *Task) Dispatch() {
	for _, worker := range workers[t.event] {
		worker.Do()
	}
}
