package main

import (
	"fmt"
)

//Tablet represents the Tablet device
type Tablet struct {
	Name string
	IsOn bool
	Volume int
	OnButton Button
	OffButton Button
	PlusVolume Button
	MinusVolume Button
}

//On turns on the Tablet
func (tab *Tablet) On() {
	tab.IsOn = true
	fmt.Println("Turning the Tablet on");
}

//Off turns off the TV
func (tab *Tablet) Off() {
	tab.IsOn = false
	fmt.Println("Turning the Tablet Off");
}

//IncreaseVolume louder the TV sound
func (tab *Tablet) IncreaseVolume() (err error) {
	err = nil
	if !tab.IsOn {
		err = ErrDeviceIsOff		
		return 
	}
	if tab.Volume >=0 && tab.Volume <100 {
		tab.Volume++
		fmt.Println("Actual ", tab.Name, " volume is ", tab.Volume)
	} else {
		fmt.Println(tab.Name, "is at maximum volume ", tab.Volume)
	}
	return
}

//DecreaseVolume turns down the TV sound
func (tab *Tablet) DecreaseVolume() (err error) {
	err = nil
	if !tab.IsOn {
		err = ErrDeviceIsOff		
		return 
	}
	if tab.Volume >0 && tab.Volume <=100 {
		tab.Volume--
		fmt.Println("Actual ", tab.Name, " volume is ", tab.Volume)
	} else {
		fmt.Println(tab.Name, " is mute. Volume ", tab.Volume)
	}
	return
}

//NewTablet return new NewTablet object
func NewTablet(name string) (tab *Tablet) {
	tab = &Tablet{Name: name}
	tab.OnButton = Button{command: &CommandOn{device: tab} }
	tab.OffButton = Button{command: &CommandOff{device: tab} }
	tab.PlusVolume = Button{command: &CommandIncreaseVolume{device: tab} }
	tab.MinusVolume = Button{command: &CommandDecreaseVolume{device: tab} }
	return tab
}
