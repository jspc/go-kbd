package main

import (
	"fmt"
	"log"

	evdev "github.com/gvalkov/golang-evdev"
)

type Event interface {
	String() string
}

func main() {
	kbdEvent, err := NewK().Lookup()
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

	for i := 0; i >= 0; i++ {
		ievent, err := dev.ReadOne()
		if err != nil {
			log.Fatal(err)
		}

		if ievent.Type == 1 && ievent.Value == 1 {
			kevent := evdev.NewKeyEvent(ievent)
			log.Println(kevent.Scancode)
		}
	}
}

func KBLog(eventID int, e Event) {
	log.Printf("[%d] -> %s", eventID, e.String())
}
