package main

import (
    "testing"
    "strings"
)

// Helper functions
func strToBool(in string) []bool {
	in = strings.ReplaceAll(in, " ", "")
	output := make([]bool, len(in))
	for i, c := range in {
		output[i] = c == '1'
	}
	return output
}

func boolToStr(in []bool) string {
	output := ""
	for _, v := range in {
		binary := ""
		if v {
			binary = "1"
		} else {
			binary = "0"
		}
		output = output + binary
	}
	return output
}

func assertSlice(actual []bool, expected []bool, t *testing.T) {
	if len(expected) != len(actual) {
		t.Error("Length is different.")
	}
	for index, value := range expected {
		if value != actual[index] {
			t.Errorf("Actual %s is different from expected %s", boolToStr(actual), boolToStr(expected))
		}
	}
}


func TestNand(t *testing.T) {
    // The only false case for nand is when both arguments are true 
    if nand(true, true) != false {
        t.Error("arg1: true, arg2: true, expected: false, got: true")
	}
	if nand(true, false) != true {
        t.Error("arg1: true, arg2: false, expected: true, got: false")
	}
	if nand(false, true) != true {
        t.Error("arg1: false, arg2: true, expected: true, got: false")
	}
	if nand(false, false) != true {
        t.Error("arg1: false, arg2: false, expected: true, got: false")
	}
}


func TestAnd(t *testing.T) {
    // The only false case for nand is when both arguments are true 
    if and(true, true) != true {
        t.Error("arg1: true, arg2: true, expected: true, got: false")
	}
	if and(true, false) != false {
        t.Error("arg1: true, arg2: false, expected: false, got: true")
	}
	if and(false, true) != false {
        t.Error("arg1: false, arg2: true, expected: false, got: true")
	}
	if and(false, false) != false {
        t.Error("arg1: false, arg2: false, expected: false, got: true")
	}
}


func TestXor(t *testing.T) {
    // The only false case for nand is when both arguments are true 
    if xor(true, true) != false {
        t.Error("arg1: true, arg2: true, expected: false, got: true")
	}
	if xor(true, false) != true {
        t.Error("arg1: true, arg2: false, expected: true, got: false")
	}
	if xor(false, true) != true {
        t.Error("arg1: false, arg2: true, expected: true, got: false")
	}
	if xor(false, false) != false {
        t.Error("arg1: false, arg2: false, expected: false, got: true")
	}
}

func TestMux(t *testing.T) {
    if mux(false, false, false) != false {
        t.Error("arg1: false, arg2: false, sel: false, expected: false, got:true")
    }
    if mux(false, true, false) != false {
        t.Error("arg1: false, arg2: true, sel: false, expected: false, got:true")
    }
    if mux(true, false, false) != true {
        t.Error("arg1: true, arg2: false, sel: false, expected: true, got: false")
    }
    if mux(true, true, false) != true {
        t.Error("arg1: true, arg2: true, sel: false, expected: true, got: false")
    }
    if mux(false, false, true) != false {
        t.Error("arg1: false, arg2: false, sel: true, expected: false, got:true")
    }
    if mux(false, true, true) != true {
        t.Error("arg1: false, arg2: true, sel: true, expected: true, got: false")
    }
    if mux(true, false, true) != false {
        t.Error("arg1: true, arg2: false, sel: true, expected: false, got: true")
    }
    if mux(true, true, true) != true {
        t.Error("arg1: true, arg2: true, sel: true, expected: true, got: false")
    }
}

func TestDmux(t *testing.T) {
    out1, out2 := dmux(true, true)
    if out1 != false || out2 != true {
        t.Error("input: true, sel: true, expected: (false, true), got: ", out1, out2)
    }
    out1, out2 = dmux(true, false)
    if out1 != true || out2 != false {
        t.Error("input: true, sel: false, expected: (true, false), got: ", out1, out2)
    }
    out1, out2 = dmux(false, true)
    if out1 != false || out2 != false {
        t.Error("input: false, sel: true, expected: (false, false), got: ", out1, out2)
    }
    out1, out2 = dmux(false, false)
    if out1 != false || out2 != false {
        t.Error("input: false, sel: false, expected: (fasle, false), got: ", out1, out2)
    }
}


func TestNotMult(t *testing.T) {
    test := []string{"0001", "0011", "0111", "1111", "1010"}
    result := []string{"1110", "1100", "1000", "0000", "0101"}

    var out []bool

    for index := range test {
        out = notMult(strToBool(test[index]))
        assertSlice(out, strToBool(result[index]), t)
    }
}

func TestAndMult(t *testing.T) {
	result := andMult(strToBool("1100"), strToBool("1010"))
	assertSlice(result, strToBool("1000"), t)
}

func TestOrMult(t *testing.T) {
	result := orMult(strToBool("1100"), strToBool("1010"))
	assertSlice(result, strToBool("1110"), t)
}

func TestMuxMult(t *testing.T) {
	result := muxMult(strToBool("1100"), strToBool("1010"), false)
	assertSlice(result, strToBool("1100"), t)
	result = muxMult(strToBool("1100"), strToBool("1010"), true)
	assertSlice(result, strToBool("1010"), t)
}

func TestDmuxMult(t *testing.T) {
	result1, result2 := dmuxMult(strToBool("11"), false)
	assertSlice(result1, strToBool("11"), t)
	assertSlice(result2, strToBool("00"), t)
	result1, result2 = dmuxMult(strToBool("11"), true)
	assertSlice(result1, strToBool("00"), t)
	assertSlice(result2, strToBool("11"), t)
}

func TestOrMultWay(t *testing.T) {
	if orMultWay(strToBool("000")) != false {
		t.Error("All false != false")
	}
	if orMultWay(strToBool("010")) != true {
		t.Error("One true != true")
	}
}

func TestMux4WayMult(t *testing.T) {
	result := mux4WayMult(strToBool("11"), strToBool("10"), strToBool("01"), strToBool("00"), strToBool("00"))
	assertSlice(result, strToBool("11"), t)
	result = mux4WayMult(strToBool("11"), strToBool("10"), strToBool("01"), strToBool("00"), strToBool("01"))
	assertSlice(result, strToBool("10"), t)
	result = mux4WayMult(strToBool("11"), strToBool("10"), strToBool("01"), strToBool("00"), strToBool("10"))
	assertSlice(result, strToBool("01"), t)
	result = mux4WayMult(strToBool("11"), strToBool("10"), strToBool("01"), strToBool("00"), strToBool("11"))
	assertSlice(result, strToBool("00"), t)
}

func TestMux8WayMult(t *testing.T) {
	result := mux8WayMult(strToBool("111"), strToBool("101"), strToBool("011"), strToBool("001"), strToBool("110"), strToBool("100"), strToBool("010"), strToBool("000"), strToBool("000"))
	assertSlice(result, strToBool("111"), t)
	result = mux8WayMult(strToBool("111"), strToBool("101"), strToBool("011"), strToBool("001"), strToBool("110"), strToBool("100"), strToBool("010"), strToBool("000"), strToBool("101"))
	assertSlice(result, strToBool("100"), t)
}

func TestDmux4Way(t *testing.T) {
	result1, result2, result3, result4 := dmux4Way(true, strToBool("00"))
	if result1 != true || result2 != false || result3 != false || result4 != false {
		t.Error("Select 1 failed")
	}
	result1, result2, result3, result4 = dmux4Way(true, strToBool("01"))
	if result1 != false || result2 != true || result3 != false || result4 != false {
		t.Error("Select 2 failed")
	}
	result1, result2, result3, result4 = dmux4Way(true, strToBool("10"))
	if result1 != false || result2 != false || result3 != true || result4 != false {
		t.Error("Select 3 failed")
	}
	result1, result2, result3, result4 = dmux4Way(true, strToBool("11"))
	if result1 != false || result2 != false || result3 != false || result4 != true {
		t.Error("Select 4 failed")
	}
}

func TestDmux8Way(t *testing.T) {
	o1, _, _, _, _, _, _, _ := dmux8Way(true, strToBool("000"))
	if o1 != true {
		t.Error("Select 1 failed")
	}
	_, o2, _, _, _, _, _, _ := dmux8Way(true, strToBool("001"))
	if o2 != true {
		t.Error("Select 2 failed")
	}
	_, _, _, _, _, o6, _, _ := dmux8Way(true, strToBool("101"))
	if o6 != true {
		t.Error("Select 5 failed")
	}
}

