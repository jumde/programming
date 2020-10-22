/*
 *  Valid Sudoku:
 *    - no duplicates in row
 *    - no duplicates in column
 *    - no duplicates in box
 */

#include <iostream>

using namespace std;

int isValidConfig(char arr[9][9]) {
  int temp = 0;
  int a[9] = {0, 0, 0, 0, 0, 0, 0, 0, 0};

  // check for valid rows
  for (int i = 0; i < 9; i++) {
    for (int j = 0; j < 9; j++) {
      if (arr[i][j] == '.') {
        continue;
      }

      temp = arr[i][j] - '1';
      if (temp < 0 || temp > 8) {
        return false;
      }

      if (a[temp] == 1) {
        return false;
      }
      a[temp] = 1;
    }
    for (int z = 0; z < 9; z++) {
      a[z] = 0;
    }
  }

  // check for valid rows
  for (int j = 0; j < 9; j++) {
    for (int i = 0; i < 9; i++) {
      if (arr[i][j] == '.') {
        continue;
      }

      temp = arr[i][j] - '1';
      if (temp < 0 || temp > 8) {
        return false;
      }

      if (a[temp] == 1) {
        return false;
      }
      a[temp] = 1;
    }
    for (int z = 0; z < 9; z++) {
      a[z] = 0;
    }
  }

  // check for valid boxes
  for (int i = 0; i <= 6; i += 3) {
    for (int j = 0; j <= 6; j += 3) {
      for (int x = i; x < i + 3; x++) {
        for (int y = j; y < j + 3; y++) {
          if (arr[x][y] == '.') {
            continue;
          }
          if (arr[x][y] < 9) {
            return false;
          }
          temp = arr[x][y] - '1';
          if (a[temp] == 1) {
            return false;
          }
          a[temp] = 1;
        }
      }
      for (int z = 0; z < 9; z++) {
        a[z] = 0;
      }
    }
  }

  return true;
}

int main() {
  char board[9][9] = {{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
                      {'4', '.', '.', '6', '9', '5', '.', '.', '.'},
                      {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
                      {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
                      {'6', '.', '.', '8', '.', '3', '.', '.', '1'},
                      {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
                      {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
                      {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
                      {'.', '.', '.', '.', '8', '.', '.', '7', '9'}};

  cout << "Is valid config: " << isValidConfig(board) << "\n";
}
