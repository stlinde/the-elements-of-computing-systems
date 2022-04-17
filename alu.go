package main

func halfAdder(arg1 bool, arg2 bool) (bool, bool) {
    carry := and(arg1, arg2)
    sum := xor(arg1, arg2)
    return carry, sum
}

func fullAdder(arg1 bool, arg2 bool, arg3 bool) (bool, bool) {
    carry_23, sum_23 := halfAdder(arg2, arg3)
    carry_123, sum_123 := halfAdder(arg1, sum_23)
    carry, sum := or(carry_23, carry_123), sum_123
    return carry, sum
}


// The add function does not check or handle overflow.
// When going throguh the code after finishing the project, add this.
func adder(arg1 []bool, arg2[]bool) []bool {
    // We need an array to store the result and a var to store the carr. 
    result := make([]bool, len(arg1))
    carry := false
    length := len(arg1) - 1

    for index := range arg1 {
        carry, result[length-index] = fullAdder(arg1[length-index], arg2[length-index], carry)
    }
    return result
}

// The inc function does not check or handle overflow.
// When going through the coe after finishing the project, add this.
func inc(arg1 []bool) []bool {
    incr := make([]bool, len(arg1))
    incr[len(incr)-1] = true
    return adder(arg1, incr)
}

/*
* The ALU that will be implemented below will take two n-bit binaries as input.
* Furthermore, six control bits will be passed to the ALU.
* These will determine what the ALU will do to the two inputs:
* Args:
*   arg1 []bool:    The first binary array input.
*   arg2 []bool:    The second binary array input.
*   zx bool:        Zero the arg1 input.
*   nx bool:        Negate the arg1 input.
*   zy bool:        Zero the arg2 input.
*   ny bool:        Negate the arg2 input.
*   f bool:         Function code: 1 for Add, 0 for And.
*   no bool:        Negate the out output.
*
* Outputs:
*   out []bool:     The output of the ALU.
*   zr bool:        True iff out=0.
*   ng bool:        True iff out <0
*/

func alu(
    arg1 []bool,
    arg2 []bool,
    zx bool,
    nx bool,
    zy bool,
    ny bool,
    f bool,
    no bool,
) ([]bool, bool, bool) {
    // Checking arg1 control bits and converting if on.
    arg1 = zeroMultControl(arg1, zx)
    arg1 = muxMult(arg1, notMult(arg1), nx) 
    
    // Checking arg2 control bits and converting if on.
    arg2 = zeroMultControl(arg1, zy)
    arg2 = muxMult(arg2, notMult(arg2), ny)

    // Running through function, 1 for Add, 0 for And.
    out := muxMult(adder(arg1, arg2), andMult(arg1, arg2), f)
    
    // Checing if output should be negated 
    out = muxMult(out, notMult(out), no)

    zr := not(orMultWay(out))
    ng := out[0]

    return out, zr, ng 
}
