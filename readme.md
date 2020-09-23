# splash

Splash is a light-weight threadpool library written in go for all generic use cases of a workerpool/threadpool.
It uses buffered channels to hold up incoming work request objects and workers.
It gets an existing worker from the workerpool and assigns the incoming task to it and then the worker performs the task.

# Features!

  - Easy to use plug and play interface
  - Low memory footprint

You can also:
  - configure the amount of log messages generated, the maximum concurrency and queue capacity.

### Tech

Built from ground-up entire using [Golang](https://golang.org/)

### Installation
`go get -t github.com/monodeepdas1215/splash`

### Examples

##### Steps
- Initialize the threadpool

```
// pool := core.NewSplashPool(requestBufferSize int, maxConcurrency int, logLevel int)
pool := core.NewSplashPool(500, 1000, core.InfoLevel)
```
Log Level can be any of the following:

`core.DebugLevel = 1 | core.InfoLevel = 2 | core.WarningLevel = 3 | core.ErrorLevel = 0`

- Create a WorkRequest struct
```
type ABC struct{
    // some fields if required
}
```
- [Mandatory] Override the functions Execute() and GetId() for the struct created. This will implement the IWorkRequest interface
```
func (abc *ABC) Execute() {
    // do your task here
}

func (abc *ABC) GetID() string {
    return "<some string that will represent the id for this struct>"
}
```
- Add the WorkRequest struct to the server
```
pool.AddWorkRequest(&ABC{})
```
License
-----

MIT