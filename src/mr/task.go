package mr

import (
	"fmt"
)

const (
	MapTask    = 0
	ReduceTask = 1
)

const (
	TaskStatusIdle    = 0
	TaskStatusRunning = 1
	TaskStatusDone    = 2
	TaskStatusError   = 3
)

type Task struct {
	Index  int
	Type   int
	Status int
	Ctx    interface{}
}

// print Task info
func (t Task) String() string {
	return fmt.Sprintf("Task{Index: %d, Type: %d, Status: %d, Ctx: %s}", t.Index, t.Type, t.Status, t.Ctx)
}
