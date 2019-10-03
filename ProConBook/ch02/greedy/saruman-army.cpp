#include <iostream>
#include <algorithm>
using namespace std;

const int N = 10;
const int R = 10;
int X[N] = {1, 7, 15, 20, 30, 50, 85, 86, 95, 106};

int main() {
  sort(X, X + N);
  int idx = 0, cnt = 0, base = 0;

  while (idx < N) {
    while (X[base + 1] - X[idx] <= R) base++;
    idx = base;
    while (X[idx + 1] - X[base] <= R) idx++;
    cnt++;
    idx++;
    // make base equal to idx for next loop
    base = idx;
    if (idx >= N) break;
  }

  printf("%d\n", cnt);
  return 0;
}