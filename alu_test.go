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

func TestAdd(t *testing.T) {
    arg1 := []bool{false, true, true, false}
    arg2 := []bool{true, false, false, true}
    expected := []bool{true, true, true, true}
    result := adder(arg1, arg2)
    testMult(expected, result, t)


    arg1 = []bool{false, false, true, false, true, true, true, true}
    arg2 = []bool{true, false, false, true, false, true, false,true}
    expected = []bool{true, true, false, false, false, true, false, false}
    result = adder(arg1, arg2)
    testMult(expected, result, t)
}

func TestIncr(t *testing.T) {
    arg1 := []bool{false, true, true, false}
    expected := []bool{false, true, true, true}
    result := inc(arg1)
    testMult(expected, result, t)

    arg1 = []bool{false, false, true, false, true, true, true, true}
    expected = []bool{false, false, true, true, false, false, false, false}
    result = inc(arg1)
    testMult(expected, result, t)
}
