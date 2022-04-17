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

