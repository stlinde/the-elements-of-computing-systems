#include <iostream>

#include "gates.h"

auto TestNand() -> void {
  if (nand(true, true) != false) {
    std::cerr << "arg1: true, arg2: true, expected: false, got: true";
  }
  if (nand(true, false) != true) {
    std::cerr << "arg1: true, arg2: false, expected: true, got: false";
  }
  if (nand(false, true) != true) {
    std::cerr << "arg1: false, arg2: true, expected: true, got: false";
  }
  if (nand(false, false) != true) {
    std::cerr << "arg1: false, arg2: false, expected: true, got: false";
  }
}

auto main() -> int {
  TestNand();
  return 0;
}
