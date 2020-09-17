package main

import (
	"fmt"
	"github.com/monodeepdas1215/splash/core"
	"time"
)

type myWorkReq struct {
}

func (mw *myWorkReq) Execute() {
	for i := 0; i < 2; i++ {
		fmt.Println("hello from myWorkReq: ", time.Now())
		time.Sleep(5 * time.Second)
	}
}

func (mw *myWorkReq) GetId() string {
	return "my-work-request"
}

type abc struct {
}

func (ab *abc) Execute() {
	for i:=0; i < 2; i ++ {
		fmt.Println("hello from abc: ", time.Now())
		time.Sleep(2 * time.Second)
	}
	time.Sleep(2 * time.Second)
}

func (ab *abc) GetId() string {
	return "abc"
}

func main() {

	pool := core.NewSplashPool(5, 10, core.DebugLevel)
	pool.StartPool()

	pool.AddWorkRequest(&abc{})
	pool.AddWorkRequest(&myWorkReq{})

	time.Sleep(2 * time.Minute)
}
