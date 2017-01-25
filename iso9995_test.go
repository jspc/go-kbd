package main

import (
    "testing"
)

func TestNewISO9995(t *testing.T) {
    isokb := NewISO9995()
    emptykb := ISO9995{}

    if isokb != emptykb {
        t.Errorf("NewISO9995() error: received %q", isokb)
    }
}

func TestKeycodeValues(t *testing.T) {
    i := ISO9995{}

    for _,tv := range []struct{
        name string
        m func()uint16
        value uint16
    }{
        {"CapsLock()", i.CapsLock, 58},
        {"LeftShift()", i.LeftShift, 42},
        {"RightShift()", i.RightShift, 54},
        {"Alt()", i.Alt, 56},
        {"AltGr()", i.AltGr, 100},
    } {
        t.Run(tv.name, func(t *testing.T) {
            if tv.m() != tv.value {
                t.Errorf("%s: expected %d, received %d", tv.name, tv.value, tv.m())
            }
        })
    }
}
