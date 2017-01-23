package main

import (
	"fmt"
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
	kbdEvent, err := keyboard.Lookup()
	if err != nil {
		log.Fatal(err)
	}

	kbd := fmt.Sprintf("/dev/input/%s", kbdEvent)

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
				case scanCode == keyboard.Mapper.LeftShift || scanCode == keyboard.Mapper.RightShift:
					keyboard.Mapper.ShiftOff()
				case scanCode == keyboard.Mapper.Alt || scanCode == keyboard.Mapper.AltGr:
					keyboard.Mapper.AltOff()
				}
			}

			if ievent.Value == 1 {
				switch {
				case scanCode == keyboard.Mapper.LeftShift || scanCode == keyboard.Mapper.RightShift:
					keyboard.Mapper.ShiftOn()
				case scanCode == keyboard.Mapper.Alt || scanCode == keyboard.Mapper.AltGr:
					keyboard.Mapper.AltOn()
				case scanCode == keyboard.Mapper.CapsLock:
					keyboard.Mapper.CapsLockFlip()
				default:
					keyboard.Mapper.Print(scanCode)
				}
			}
		}
	}
}
