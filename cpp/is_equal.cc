#include <iostream>

using namespace std;

bool isEqual(int a, int b) {
  return !(a ^ b);
}

int main() {
  cout << "isEqual (8, 8): " << isEqual(8, 8) << "\n";
  cout << "isEqual (7, 6): " << isEqual(7, 6) << "\n";
  cout << "isEqual (1, 0): " << isEqual(1, 0) << "\n";
  cout << "isEqual (0, 0): " << isEqual(0, 0) << "\n";
  return 0;
}
