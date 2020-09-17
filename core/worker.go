package core

import (
	"fmt"
	"time"
)

type workerThread struct {
	id			int
	waitChannel	chan IWorkRequest
}

func NewWorkerThread(id int) *workerThread {
	obj := &workerThread{
		id:          id,
		waitChannel: make(chan IWorkRequest),
	}

	obj.run()
	return obj
}

func (worker *workerThread) run() {

	AppLogger.logger.Infof("%s started running... waiting for incoming work requests", worker.GetWorkerDetails())

	go func() {
		for {
			select {
			case incomingWorkReq := <- worker.waitChannel:
				t1 := time.Now()
				incomingWorkReq.Execute()
				diff := time.Now().Sub(t1).Seconds()
				AppLogger.logger.Infof("%s finished work-req : %s in %f secs", worker.GetWorkerDetails(), incomingWorkReq.GetId(), diff)
			}
		}
	}()
}

func (worker *workerThread) GetWorkerDetails() string {
	return fmt.Sprintf("Worker(id = %d)", worker.id)
}

func (worker *workerThread) AssignWork(workReq IWorkRequest) {
	worker.waitChannel <- workReq
}