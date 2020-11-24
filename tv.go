package main

import (
	"fmt"
)

//TV represents the TV device
type TV struct {
	IsOn bool
	Volume int
	OnButton Button
	OffButton Button
	PlusVolume Button
	MinusVolume Button
}

//On turns on the TV
func (tv *TV) On() {
	tv.IsOn = true
	fmt.Println("Turning the TV on");
}

//Off turns off the TV
func (tv *TV) Off() {
	tv.IsOn = false
	fmt.Println("Turning the TV Off");
}

//IncreaseVolume louder the TV sound
func (tv *TV) IncreaseVolume() (err error) {
	err = nil
	if !tv.IsOn {
		err = ErrDeviceIsOff		
		return 
	}
	if tv.Volume >=0 && tv.Volume <100 {
		tv.Volume++
		fmt.Println("Actual TV volume is ", tv.Volume)
	} else {
		fmt.Println("TV is at maximum volume ", tv.Volume)
	}
	return
}

//DecreaseVolume turns down the TV sound
func (tv *TV) DecreaseVolume() (err error) {
	err = nil
	if !tv.IsOn {
		err = ErrDeviceIsOff		
		return 
	}
	if tv.Volume >0 && tv.Volume <=100 {
		tv.Volume--
		fmt.Println("Actual TV volume is ", tv.Volume)
	} else {
		fmt.Println("TV is mute. Volume ", tv.Volume)
	}
	return
}

//NewTV return new TV object
func NewTV() (tv *TV) {
	tv = &TV{}
	tv.OnButton = Button{command: &CommandOn{device: tv} }
	tv.OffButton = Button{command: &CommandOff{device: tv} }
	tv.PlusVolume = Button{command: &CommandIncreaseVolume{device: tv} }
	tv.MinusVolume = Button{command: &CommandDecreaseVolume{device: tv} }
	return tv
}
