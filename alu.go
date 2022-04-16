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
