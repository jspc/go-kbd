package main

import (
    "testing"
)

func TestNewK(t *testing.T) {
    k := NewK()
    if k.Proc != "/proc/bus/input/devices" {
        t.Errorf("NewK() error: expected '/proc/bus/input/devices', received %q", k.Proc)
    }
}

func TestKBLookup(t *testing.T) {
    k := K{"./fixtures/devices"}

    kbd,err := k.Lookup()
    if err != nil {
        t.Errorf("Lookup() error: %q", err)
    }

    if kbd != "event4" {
        t.Errorf("Lookup() error: expected 'event4', received %q", kbd)
    }
}
