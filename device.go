package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// K represents a keyboard
type K struct {
	// Proc contains a filename where one might find keyboard details. In proc, usually
	// - this differs for testing and, potentially, different OSes
	Proc string

	Caps     bool
	Shifted  bool
	Alted    bool
	Mappings []Mapping

	LeftShift, RightShift, Alt, AltGr, CapsLock uint16
}

var (
	bitMask           = "120013"
	procDefaultString = "/proc/bus/input/devices"

	i int
	l string

	mm MapManager
)

func NewK(t string) (k K) {
	switch t {
	case "iso9995":
		mm = NewISO9995()
	}

	k.Proc = procDefaultString

	// TODO: How would one ensure these *are* false at start up?
	k.Caps = false
	k.Shifted = false
	k.Alted = false

	k.Mappings = mm.Mappings()
	k.LeftShift = mm.LeftShift()
	k.RightShift = mm.RightShift()
	k.Alt = mm.Alt()
	k.AltGr = mm.AltGr()
	k.CapsLock = mm.CapsLock()

	return
}

func (k K) Path() (d string, err error) {
	var proc []byte

	proc, err = ioutil.ReadFile(k.Proc)
	if err != nil {
		return
	}

	procSplit := strings.Split(string(proc), "\n")
	for i, l = range procSplit {
		if strings.Contains(l, fmt.Sprintf("EV=%s", bitMask)) {
			break
		}
	}

	for c := i; c > 0; c-- {
		l = procSplit[c]
		if l == "" {
			err = fmt.Errorf("Could not find a device ID for EV=%s", bitMask)
			return
		}
		if strings.Contains(l, "Handlers=") {
			r, _ := regexp.Compile("event[0-9]+")
			d = r.FindStringSubmatch(l)[0]
			break
		}
	}

	d = fmt.Sprintf("/dev/input/%s", d)
	return
}
