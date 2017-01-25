package main

import (
	"testing"
)

func TestNewK(t *testing.T) {
	k := NewK("iso9995")
	if k.Proc != "/proc/bus/input/devices" {
		t.Errorf("NewK() error: expected '/proc/bus/input/devices', received %q", k.Proc)
	}
}

func TestKBLookup(t *testing.T) {
	k := K{"./fixtures/devices", false, false, false, []Mapping{}, 0, 0, 0, 0, 0}

	kbd, err := k.Path()
	if err != nil {
		t.Errorf("Lookup() error: %q", err)
	}

	if kbd != "/dev/input/event4" {
		t.Errorf("Lookup() error: expected 'event4', received %q", kbd)
	}
}
