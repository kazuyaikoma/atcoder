#include <iostream>
using namespace std;

bool dfs(int i, int sum);
const int n = 4;
int a[n] = {1, 2, 4, 7};
int k = 13;

int main() {
  if (dfs(0, 0)) printf("Yes\n");
  else printf("No\n");
  return 0;
}

bool dfs(int i, int sum) {
  if (i == n) return sum == k;

  // a[i]を使う場合
  if (dfs(i + 1, sum)) return true;
  // a[i]を使わない場合
  if (dfs(i + 1, sum + a[i])) return true;

  return false;
}
