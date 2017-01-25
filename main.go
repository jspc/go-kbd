package main

import (
	"log"

	evdev "github.com/gvalkov/golang-evdev"
)

type Event interface {
	String() string
}

var (
	scanCode uint16
)

func main() {
	keyboard := NewK("iso9995")
	kbd, err := keyboard.Path()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(kbd)

	dev, err := evdev.Open(kbd)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(dev)

	for {
		ievent, err := dev.ReadOne()
		if err != nil {
			log.Fatal(err)
		}

		if ievent.Type == 1 {
			kevent := evdev.NewKeyEvent(ievent)
			scanCode = kevent.Scancode

			if ievent.Value == 0 {
				switch {
				case scanCode == keyboard.LeftShift || scanCode == keyboard.RightShift:
					keyboard.ShiftOff()
				case scanCode == keyboard.Alt || scanCode == keyboard.AltGr:
					keyboard.AltOff()
				}
			}

			if ievent.Value == 1 {
				switch {
				case scanCode == keyboard.LeftShift || scanCode == keyboard.RightShift:
					keyboard.ShiftOn()
				case scanCode == keyboard.Alt || scanCode == keyboard.AltGr:
					keyboard.AltOn()
				case scanCode == keyboard.CapsLock:
					keyboard.CapsLockFlip()
				default:
					keyboard.Print(scanCode)
				}
			}
		}
	}
}
