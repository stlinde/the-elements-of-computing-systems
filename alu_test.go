package main

import "testing"

func TestHalfAdder(t *testing.T) {
    out1, out2 := halfAdder(false, false)
    if out1 != false || out2 != false {
        t.Error("arg1: false, arg2: false, expected: (false, false), got: ", out1, out2)
    }
        
    out1, out2 = halfAdder(false, true)
    if out1 != false || out2 != true {
        t.Error("arg1: false, arg2: true, expected: (false, true), got: ", out1, out2)
    }

    out1, out2 = halfAdder(true, false)
    if out1 != false || out2 != true {
        t.Error("arg1: true, arg2: false, expected: (false, true), got: ", out1, out2)
    }

    out1, out2 = halfAdder(true, true)
    if out1 != true || out2 != false {
        t.Error("arg1: true, arg2: true, expected: (true, false), got: ", out1, out2)
    }
}

func TestFullAdder(t *testing.T) {
    out1, out2 := fullAdder(false, false, false)
    if out1 != false || out2 != false {
        t.Error("arg1: false, arg2: false, arg3: false, expected: (false, false), got: ", out1, out2)
    }
        
    out1, out2 = fullAdder(false, false, true)
    if out1 != false || out2 != true {
        t.Error("arg1: false, arg2: false, arg3: true expected: (false, true), got: ", out1, out2)
    }

    out1, out2 = fullAdder(false, true, false)
    if out1 != false || out2 != true {
        t.Error("arg1: false, arg2: true, arg3: false, expected: (false, true), got: ", out1, out2)
    }

    out1, out2 = fullAdder(false, true, true)
    if out1 != true || out2 != false {
        t.Error("arg1: false, arg2: true, arg3: true, expected: (true, false), got: ", out1, out2)
    }
    out1, out2 = fullAdder(true, false, false)
    if out1 != false || out2 != true {
        t.Error("arg1: true, arg2: false, arg3: false, expected: (false, true), got: ", out1, out2)
    }
        
    out1, out2 = fullAdder(true, false, true)
    if out1 != true || out2 != false {
        t.Error("arg1: true, arg2: false, arg3: true expected: (true, false), got: ", out1, out2)
    }

    out1, out2 = fullAdder(true, true, false)
    if out1 != true || out2 != false {
        t.Error("arg1: true, arg2: true, arg3: false, expected: (true, false), got: ", out1, out2)
    }

    out1, out2 = fullAdder(true, true, true)
    if out1 != true || out2 != true {
        t.Error("arg1: true, arg2: true, arg3: true, expected: (true, true), got: ", out1, out2)
    }
}
