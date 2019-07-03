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
	fmt.Println("phone start working...")
}

func (p Phone) Stop() {
	fmt.Println("phone stop working...")
}

type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println("Camera start working...")
}

func (c Camera) Stop() {
	fmt.Println("Camera stop working...")
}

func main() {
	var usbArr [3]Usb
	usbArr[0] = Phone{"OnePlus"}
	usbArr[1] = Phone{"Vivo"}
	usbArr[2] = Camera{"Niko"}
	fmt.Println(usbArr)
}
