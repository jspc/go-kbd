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
	Proc   string
	Mapper Mapper
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
	k.Mapper = NewMapper(mm)

	return
}

func (k K) Lookup() (d string, err error) {
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
	return
}
