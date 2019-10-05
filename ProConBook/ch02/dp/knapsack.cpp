#include <iostream>
#include <vector>
#include <cstring>
using namespace std;

const int N = 4;
const int W = 5;

int w[N] = {2, 1, 3, 2};
int v[N] = {3, 2, 4, 2};

// memo
int dp[N+1][W+1];
// interface
int calc(int idx, int limit_w);

int main() {
  memset(dp, -1, sizeof(dp));
  printf("%d\n", calc(0, W));
  return 0;
}

int calc(int idx, int limit_w) {
  if (dp[idx][limit_w] >= 0) {
    return dp[idx][limit_w];
  }

  int ret = 0;
  if (idx == N) {
    ret = 0;
  } else if (limit_w < w[idx]) {
    ret = calc(idx + 1, limit_w);
  } else {
    ret = max(
      calc(idx + 1, limit_w),
      calc(idx + 1, limit_w - w[idx]) + v[idx]
    );
  }
  return dp[idx][limit_w] = ret;
}