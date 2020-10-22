/*
 *  - binary string a: 1001
 *  - binary string b: 1100
 *  - sum: 10101
 */

#include <iostream>

using namespace std;

string generateSumString(string a, string b) {
  string sum_string = "";

  uint length_a = a.length();
  uint length_b = b.length();

  uint current_a = 0; // track current a
  uint current_b = 0; // track current b 
  uint current_c = 0;
  uint carryover = 0;

  int i = length_a - 1;
  int j = length_b - 1;

  char c = '0';

  while (i >= 0 && j >= 0) {
    current_a = int(a[i]) - int('0');
    current_b = int(b[j]) - int('0');

    current_c = current_a ^ current_b ^ carryover;
    carryover = (current_a + current_b + carryover) / 2;

    c = char(int('0') + current_c);
    sum_string += c;

    i--;
    j--;
  }

  while (i >= 0) {
    current_a = int(a[i]) - int('0');
    current_c = current_a ^ carryover;
    carryover = (current_a + carryover) / 2;
    c = char(int('0') + current_c);
    sum_string += c;
    i--;
  }

  while (j >= 0) {
    current_b = int(b[i]) - int('0');
    current_c = current_b ^ carryover;
    carryover = (current_b + carryover) / 2;
    c = char(int('0') + current_c);
    sum_string += c;
    j--;
  }

  if (carryover) { 
    c = char(int('0') + carryover);
    sum_string += c;
  }

  std::reverse(sum_string.begin(), sum_string.end());
  return sum_string;
}

int main() {
  string sum_string = generateSumString("1100", "101");
  cout << "Sum String: " << sum_string << "\n";
}
