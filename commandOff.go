package main

//CommandOff represents the Device's turn off command
type CommandOff struct {
	device Device
}

func (c *CommandOff) execute() {
	c.device.Off()
}