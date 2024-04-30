package mr

import (
	"fmt"
)

type WCTaskCtx struct {
	FileName  string
	FileIndex int
}

func (w WCTaskCtx) String() string {
	return fmt.Sprintf("{FileName: %s, FileIndex: %d}", w.FileName, w.FileIndex)
}
