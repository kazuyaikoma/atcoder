#include <iostream>
#include <vector>
#include <cstring>
using namespace std;

const int N = 4;
const int W = 5;
typedef pair<int, int> Item;
// 前者が重さ、後者が価値
vector<Item> items = {{2, 3}, {1, 2}, {3, 4}, {2, 2}};
// メモ
int dp[N+1][W+1];
// interface
int calc(int idx, int W);

int main() {
  for (int i = N - 1; i >= 0; i--) {
    for (int w = 0; w <= W; w++) {
      if (items[i].first > w) {
        dp[i][w] = dp[i+1][w];
      } else {
        dp[i][w] = max(
          dp[i+1][w],
          dp[i+1][w - items[i].first] + items[i].second
        );
      }
    }
  }

  printf("%d\n", dp[0][W]);
  return 0;
}