package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	name string
}

func (p Phone) Start() {
	fmt.Println("Phone start working......")
}

func (p Phone) Stop() {
	fmt.Println("Phone stop working......")
}

func (p Phone) Call() {
	fmt.Println("Phone is calling now.....")
}

type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println("Camera start working.....")
}

func (c Camera) Stop() {
	fmt.Println("Camera stop working.....")
}

type Computer struct {
}

func (computer Computer) Working(usb Usb) {
	usb.Start()
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	var usbArr [3]Usb
	usbArr[0] = Phone{"OnePlus"}
	usbArr[1] = Phone{"Vivo"}
	usbArr[2] = Camera{"Niko"}
	// fmt.Println(usbArr)
	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}
}
