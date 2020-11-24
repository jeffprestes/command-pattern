package main

import "fmt"

func main()  {
	tv := NewTV()
	iPad := NewTablet("iPad 10th")
	tv.OnButton.Press()
	iPad.OnButton.Press()
	tv.OffButton.Press()
	fmt.Println("main louder the TV")
	tv.PlusVolume.Press()
	tv.OnButton.Press()
	iPad.MinusVolume.Press()
	tv.PlusVolume.Press()
	tv.PlusVolume.Press()
	tv.PlusVolume.Press()
	tv.OffButton.Press()
	for i:=0; i<=20; i++ {
		iPad.PlusVolume.Press()
	}
	iPad.OffButton.Press()
}