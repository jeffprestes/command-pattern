package main

import "fmt"

//CommandIncreaseVolume represents the Device's increase volume command
type CommandIncreaseVolume struct {
	device Device
}

func (c *CommandIncreaseVolume) execute() {
	err := c.device.IncreaseVolume()
	if err != nil {
		fmt.Println("Could not execute the command. Error: ", err.Error())
	}
}