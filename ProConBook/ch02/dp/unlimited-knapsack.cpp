#include <iostream>
#include <vector>
using namespace std;

const int N = 3;
const int Weight = 7;
int W[N] = {3, 4, 2};
int V[N] = {4, 5, 3};

int dp[N+1][Weight+1];

int calc(int idx, int remain_w);

int main() {
  memset(dp, -1, sizeof(dp));
  printf("%d\n", calc(0, Weight));
  return 0;
}

int calc(int idx, int remian_w) {
  if (dp[idx][remian_w] >= 0) {
    return dp[idx][remian_w];
  }

  int ret;
  if (idx >= N) {
    ret = 0;
  } else if (W[idx] > remian_w) {
    ret = calc(idx+1, remian_w);
  } else {
    vector<int> vec = {};
    for (int c=0; c<=remian_w/W[idx]; c++) {
      vec.push_back(
        calc(idx+1, remian_w-(W[idx]*c)) + V[idx]*c
      );
    }
    ret = *max_element(vec.begin(), vec.end());
  }
  return dp[idx][remian_w] = ret;
}