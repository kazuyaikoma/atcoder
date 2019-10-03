#include <iostream>
using namespace std;

const int N = 6;
char S[N + 1] = "ACDBCB";

int main() {
  // はじめと終わりのindex
  int s_idx = 0, e_idx = N - 1;

  while (s_idx <= e_idx) {
    bool left = true;

    for (int i = 0; s_idx + i < e_idx; i++) {
      if (S[s_idx + i] < S[e_idx - i]) {
        left = true;
        break;
      } else if (S[s_idx + i] > S[e_idx - i]) {
        left = false;
        break;
      }
    }

    if (left) { 
      putchar(S[s_idx++]);
    } else {
      putchar(S[e_idx--]);
    }
  }
  
  putchar('\n');
  return 0;
}