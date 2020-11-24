package main

import "fmt"

//CommandDecreaseVolume represents the Device's decrease volume command
type CommandDecreaseVolume struct {
	device Device
}

func (c *CommandDecreaseVolume) execute() {
	err := c.device.DecreaseVolume()
	if err != nil {
		fmt.Println("Could not execute the command. Error: ", err.Error())
	}
}