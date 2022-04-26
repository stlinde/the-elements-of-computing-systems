// gates.cc
#include <vector>
#include "gates.h"

auto nand(bool arg1, bool arg2) -> bool{
  if ((arg1 == true) && (arg2 == true)) {
    return false;
  } else {
    return true; 
  }
}

auto not(bool arg1) -> bool {
  return nand(arg1, arg1);
}

auto and(bool arg1, bool arg2) -> bool {
  return not(nand(arg1, arg2));
}

auto or(bool arg1) -> {
  return nand(nand(arg1, arg1), nand(arg2, arg2));k
}

auto xor(bool arg1, bool arg2) -> bool {
  bool temp1 = nand(arg1, nand(arg1, arg2));
  bool temp2 = nand(arg2, nand(arg1, arg2));
  return nand(temp1, temp2);
}

auto mux(bool arg1, bool arg2, bool sel) -> bool {
  bool temp1 = and(arg1, not(sel));
  bool temp2 = and(arg2, sel);
  return or(temp1, temp2);
}

auto dmux(bool input, bool sel) -> bool {
  bool temp = not(sel);
  bool output1 = and(input, intermediate);
  bool output2 = and(input, sel);
  return output1, output2;
}

auto notMult(std::vector<bool> arg1) -> std::vector<bool> {
  std::vector<bool> output;
  for (int i = 0; i < arg1.size(); i++) {
    output[i] = not(arg1[i]);
  }
  return output;
}
