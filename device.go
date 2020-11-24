package main

import "errors"

//ErrDeviceIsOff an error to be return when a command is called but cannot be executed because the device is off
var ErrDeviceIsOff = errors.New("Could not execute this command because the TV is off")

//Device basic operations
type Device interface {
	On()
	Off()
	IncreaseVolume() (error)
	DecreaseVolume() (error)
}