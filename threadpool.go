package main

import (
	"fmt"
	"sync"
)

// worker pool struct to hold the worker queues and incoming work requests in buffered channels
type Splash struct {

	// buffered channel to hold all workers
	workerQ 				chan *workerThread

	// buffered channel to hold all the incoming work request
	workRequestsQ 			chan IWorkRequest
}


var initThreadPool *Splash
var singleton sync.Once


// singleton implementation to ensure only a single object of the worker pool in memory
func NewSplashPool(requestsBufferSize, maxConcurrency, loglevel int) *Splash {

	singleton.Do(func() {

		// initializing application logger
		InitLogger(loglevel, false)

		initThreadPool = &Splash{
			workerQ:		make(chan *workerThread, maxConcurrency),
			workRequestsQ:	make(chan IWorkRequest, requestsBufferSize),
		}

		// filling up the worker queue with workers
		for i := 0; i < cap(initThreadPool.workerQ); i ++ {
			worker := NewWorkerThread(i+1)
			initThreadPool.workerQ <- worker
		}
	})

	return initThreadPool
}

// adds a work request into the worker pool
func (splash *Splash) AddWorkRequest(workReq IWorkRequest) {
	splash.workRequestsQ <- workReq
}

// starts the loop to wait for incoming work requests
func (splash *Splash) StartPool() {

	AppLogger.logger.Infoln("starting splash thread-pool...")

	go func() {
		for {
			select {
			case incomingWorkReq := <-splash.workRequestsQ:

				worker := <- splash.workerQ  // get an available worker from the workerQ
				worker.AssignWork(incomingWorkReq)  // assign the incoming work to it
				splash.workerQ <- worker	// append it back to Q
			}
		}
	}()
}

func (splash *Splash) GetDetails() string {
	return fmt.Sprintf("Splash WorkerPool(maxWorkReqCap: %d, maxConcurrency: %d)", cap(splash.workRequestsQ), cap(splash.workerQ))
}