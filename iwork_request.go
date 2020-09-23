package main

type IWorkRequest interface {
	Execute()
	GetId() string
}