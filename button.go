package main

//Button represents the device button
type Button struct {
	command Command
} 

//Press represents press button action
func (button *Button) Press() {
	button.command.execute()
}
