package main

import "testing"

func testMult(expected []bool, result []bool, t *testing.T) {
    if len(expected) != len(result) {
        t.Errorf("Length error. Expected: %d, got: %d", len(expected), len(result))
    }
    for index := range expected {
        if expected[index] != result[index] {
            t.Errorf("Value wrong at index: %d, expected: %t, got: %t",
                index, expected[index], result[index])
        }
    }
}
