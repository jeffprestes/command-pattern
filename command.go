package main

//Command to be executed by device
type Command interface {
	execute()
}