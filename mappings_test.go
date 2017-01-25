package main

import (
    "testing"
)

type TestKB struct {}
func (t TestKB)Mappings() []Mapping {
    return []Mapping{
        {"a", "A", "A", "A"},
        {"1", "1", "!", "1"},
        {},
    }
}
func (t TestKB)CapsLock()uint16{return 100}
func (t TestKB)LeftShift()uint16{return 101}
func (t TestKB)RightShift()uint16{return 102}
func (t TestKB)Alt()uint16{return 103}
func (t TestKB)AltGr()uint16{return 104}

func TestMappingEmpty(t *testing.T) {
    kb := TestKB{}

    r := kb.Mappings()[2].Empty()
    if !r {
        t.Errorf("Empty(): expected true, received %q", r)
    }
}

func TestShiftKey(t *testing.T) {
    kb := TestKB{}

    t.Run("ShiftOn()", func(t *testing.T) {
        t.Run("When un-shifted", func(t *testing.T) {
            k := K{"", false, false, false, kb.Mappings(), 101, 102, 103, 104, 101}
            k.ShiftOn()

            if !k.Shifted {
                t.Errorf("k.Shifted: expected false, received %q", k.Shifted)
            }
        })

        t.Run("When shifted", func(t *testing.T) {
            k := K{"", false, true, false, kb.Mappings(), 101, 102, 103, 104, 101}
            k.ShiftOn()

            if !k.Shifted {
                t.Errorf("k.Shifted: expected false, received %q", k.Shifted)
            }
        })
    })

    t.Run("ShiftOff()", func(t *testing.T) {
        t.Run("When un-shifted", func(t *testing.T) {
            k := K{"", false, true, false, kb.Mappings(), 101, 102, 103, 104, 101}
            k.ShiftOff()

            if k.Shifted {
                t.Errorf("k.Shifted: expected true, received %q", k.Shifted)
            }
        })

        t.Run("When shifted", func(t *testing.T) {
            k := K{"", false, false, false, kb.Mappings(), 101, 102, 103, 104, 101}
            k.ShiftOff()

            if k.Shifted {
                t.Errorf("k.Shifted: expected true, received %q", k.Shifted)
            }
        })
    })
}

func TestAltKey(t *testing.T) {
    kb := TestKB{}

    t.Run("AltOn()", func(t *testing.T) {
        t.Run("When un-alted", func(t *testing.T) {
            k := K{"", false, false, false, kb.Mappings(), 101, 102, 103, 104, 101}
            k.AltOn()

            if !k.Alted {
                t.Errorf("k.Alted: expected true, received %q", k.Alted)
            }
        })

        t.Run("When alted", func(t *testing.T) {
            k := K{"", false, false, true, kb.Mappings(), 101, 102, 103, 104, 101}
            k.AltOn()

            if !k.Alted {
                t.Errorf("k.Alted: expected true, received %q", k.Alted)
            }
        })
    })

    t.Run("AltOff()", func(t *testing.T) {
        t.Run("When un-alted", func(t *testing.T) {
            k := K{"", false, false, true, kb.Mappings(), 101, 102, 103, 104, 101}
            k.AltOff()

            if k.Alted {
                t.Errorf("k.Alted: expected false, received %q", k.Alted)
            }
        })

        t.Run("When alted", func(t *testing.T) {
            k := K{"", false, false, false, kb.Mappings(), 101, 102, 103, 104, 101}
            k.AltOff()

            if k.Alted {
                t.Errorf("k.Alted: expected true, received %q", k.Alted)
            }
        })
    })
}

func TestCapsLockKey(t *testing.T) {
    kb := TestKB{}

    t.Run("CapsLockFlip()", func(t *testing.T) {
        t.Run("When off", func(t *testing.T) {
            k := K{"", false, false, false, kb.Mappings(), 101, 102, 103, 104, 101}
            k.CapsLockFlip()

            if !k.Caps {
                t.Errorf("k.Caps: expected false, received %q", k.Caps)
            }
        })

        t.Run("When on", func(t *testing.T) {
            k := K{"", true, false, false, kb.Mappings(), 101, 102, 103, 104, 101}
            k.CapsLockFlip()

            if k.Caps {
                t.Errorf("k.Caps: expected true, received %q", k.Caps)
            }
        })
    })
}
