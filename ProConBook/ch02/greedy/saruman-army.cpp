#include <iostream>
#include <algorithm>
using namespace std;

const int N = 6;
const int R = 10;
int X[N] = {1, 7, 15, 20, 30, 50};

int main() {
  sort(X, X + N);
  int idx = 0, cnt = 0;

  while (idx < N) {
    int s = X[idx++];
    while (idx < N && X[idx] - s <= R) idx++;

    // new point marked black
    int new_base = X[idx - 1];
    while (idx < N && X[idx] - new_base <= R) idx++;
    cnt++;
  }

  printf("%d\n", cnt);
  return 0;
}