#include <iostream>
#include <vector>
#include <cstring>
using namespace std;

const int N = 4;
int limit_weight = 5;
typedef pair<int, int> Item;
// 前者が重さ、後者が価値
vector<Item> items = {{2, 3}, {1, 2}, {3, 4}, {2, 2}};
// メモ
int dp[N][N];
// interface
int calc(int idx, int limit_weight);

int main() {
  memset(dp, -1, sizeof(dp));
  printf("%d\n", calc(0, limit_weight));
  return 0;
}

// メモ化して全探索
int calc(int idx, int remaining_w) {
  int ret;
  if (idx == N - 1) {
    ret = 0;
  } else if (items[idx].first > remaining_w) {
    ret = calc(idx+1, remaining_w);
  } else {
    // items[idx]を使う場合・使わない場合
    ret = max(
      calc(idx+1, remaining_w),
      calc(idx+1, remaining_w - items[idx].first) + items[idx].second
    );
  }
  return ret;
}

