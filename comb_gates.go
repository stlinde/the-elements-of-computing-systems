package main

/* Chapter 1 - Boolean Logic
 * As stated in the text, all logic gates can be constructed from the Nand
 * gate.
 * The implementation here will use the nand gate as the only building block.
 */

func nand(arg1 bool, arg2 bool) bool {
    if (arg1 == true) && (arg2 == true) {
        return false
    } else {
        return true
    }
}

func not(arg1 bool) bool {
    // The not gate can be constructed by using the input twice in the nand
    return nand(arg1, arg1)
}

func and(arg1 bool, arg2 bool) bool {
    // The And gate can be constructed using the Not gate and the Nand gate
    return not(nand(arg1, arg2))
}

func or(arg1 bool, arg2 bool) bool {
    // The Or gate can be constructed from three nand gates. 
    return nand(nand(arg1, arg1), nand(arg2, arg2))
}

func xor(arg1 bool, arg2 bool) bool {
    intermediate1 := nand(arg1, nand(arg1, arg2))
    intermediate2 := nand(arg2, nand(arg1, arg2))
    return nand(intermediate1, intermediate2)
}

func mux(arg1 bool, arg2 bool, sel bool) bool {
    intermediate1 := and(arg1, not(sel))
    intermediate2 := and(arg2, sel)
    return or(intermediate1, intermediate2)
}

func dmux(input bool, sel bool) (bool, bool) {
    intermediate := not(sel)
    output1 := and(input, intermediate)
    output2 := and(input, sel)
    return output1, output2
}

func notMult(arg1 []bool) []bool {
    output := make([]bool, len(arg1))
    for index, value := range arg1 {
        output[index] = not(value) 
    }
    return output
}

func andMult(arg1 []bool, arg2 []bool) []bool {
    output := make([]bool, len(arg1))
    for index := range arg1 {
        output[index] = and(arg1[index], arg2[index])
    }
    return output
}

func orMult(arg1 []bool, arg2 []bool) []bool {
    output := make([]bool, len(arg1))
    for index := range arg1 {
        output[index] = or(arg1[index], arg2[index])
    }
    return output
}


func muxMult(arg1 []bool, arg2 []bool, sel bool) []bool {
    output := make([]bool, len(arg1))
    for index := range arg1 {
        output[index] = mux(arg1[index], arg2[index], sel)
    }
    return output
}

func dmuxMult(input []bool, sel bool) ([]bool, []bool) {
	result1 := make([]bool, len(input))
	result2 := make([]bool, len(input))
	for index, value := range input {
		r1, r2 := dmux(value, sel)
		result1[index] = r1
		result2[index] = r2
	}
	return result1, result2
}

func orMultWay(arg1 []bool) bool {
	result := false
	for _, value := range arg1 {
		result = or(result, value)
	}
	return result
}

func mux4WayMult(arg1 []bool, arg2 []bool, arg3 []bool, arg4 []bool, sel []bool) []bool {
	result1 := muxMult(arg1, arg2, sel[1])
	result2 := muxMult(arg3, arg4, sel[1])
	return muxMult(result1, result2, sel[0])
}

func mux8WayMult(arg1 []bool, 
                 arg2 []bool, 
                 arg3 []bool, 
                 arg4 []bool,
                 arg5 []bool, 
                 arg6 []bool, 
                 arg7 []bool, 
                 arg8 []bool, 
                 sel []bool) []bool {
	result1 := mux4WayMult(arg1, arg2, arg3, arg4, sel[1:])
	result2 := mux4WayMult(arg5, arg6, arg7, arg8, sel[1:])
	return muxMult(result1, result2, sel[0])
}

func dmux4Way(input bool, sel []bool) (bool, bool, bool, bool) {
	result1, result2 := dmux(input, sel[0])
	result3, result4 := dmux(result1, sel[1])
	result5, result6 := dmux(result2, sel[1])
	return result3, result4, result5, result6
}

func dmux8Way(input bool, sel []bool) (bool, bool, bool, bool, bool, bool, bool, bool) {
	r1, r2, r3, r4 := dmux4Way(input, sel)
	o1, o2 := dmux(r1, sel[2])
	o3, o4 := dmux(r2, sel[2])
	o5, o6 := dmux(r3, sel[2])
	o7, o8 := dmux(r4, sel[2])
	return o1, o2, o3, o4, o5, o6, o7, o8
} 
  
func zeroMultControl(arg1 []bool, control bool) []bool {
    // The zeroMultControl is used as the zx and zy control bits.
    // If control is 1, it returns an array of 0's (false).
    zero_array := make([]bool, len(arg1))
    return muxMult(arg1, zero_array, control)
}
