package main

//CommandOn represents the Device's turn on Command
type CommandOn struct {
	device Device
}

func (c *CommandOn) execute() {
	c.device.On()
}