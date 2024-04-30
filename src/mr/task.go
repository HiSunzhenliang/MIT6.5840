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
	taskType := "MapTask"
	if t.Type == ReduceTask {
		taskType = "ReduceTask"
	}
	return fmt.Sprintf("Task{Index: %d, Type: %s, Status: %d, Ctx: %s}", t.Index, taskType, t.Status, t.Ctx)
}
