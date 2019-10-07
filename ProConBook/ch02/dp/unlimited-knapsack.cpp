#include <iostream>
#include <vector>
using namespace std;

const int N = 3;
const int Weight = 7;

int W[N] = {3, 4, 2};
int V[N] = {4, 5, 3};

int dp[N+1][Weight+1];

int calc(int i, int remaining_w);

int main() {
  memset(dp, -1, sizeof(dp));
  int out = calc(0, Weight);
  cout << out << "\n"s;
}

int calc(int i, int remaining_w) {
  if (dp[i][remaining_w] >= 0) return dp[i][remaining_w];

  int ret;
  if (i >= N) ret = 0;
  else if (remaining_w < W[i]) {
    ret = calc(i+1, remaining_w);
  } else {
    vector<int> vec;
    for (int k=0; k<=remaining_w/W[i]; k++) {
      vec.push_back(calc(i+1, remaining_w - W[i] * k) + V[i] * k);
    }
    ret = *max_element(vec.begin(), vec.end());
  }
  return ret;
}