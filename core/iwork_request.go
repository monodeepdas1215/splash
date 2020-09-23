package core

type IWorkRequest interface {
	Execute()
	GetId() string
}