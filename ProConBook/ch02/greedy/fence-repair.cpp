#include <iostream>
#include <algorithm>
#include <vector>
#include <numeric>
using namespace std;

const int N = 3;
vector<int> L = {8, 5, 8};
int cost = 0;

int main() {
  sort(L.begin(), L.end());
  while (L.size()) {
    int n = L.back();
    L.pop_back();
    int sum = accumulate(L.begin(), L.end(), 0);
    if (L.size() <= 0) break;
    cost += n + sum;
  }
  printf("%d\n", cost);
  return 0;
}