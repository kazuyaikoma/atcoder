#include <iostream>
#include <vector>
using namespace std;

const int N = 3;
const int Weight = 7;
int W[N] = {3, 4, 2};
int V[N] = {4, 5, 3};

int dp[N+1][Weight+1];

void calc(void);

int main() {
  calc();
  printf("%d\n", dp[N][Weight]);
  return 0;
}

void calc() {
  for (int i=0; i<N; i++) {
    for (int j=0; j<=Weight; j++) {
      if (j < W[i]) {
        dp[i+1][j] = dp[i][j];
      } else {
        dp[i+1][j] = max(
          dp[i][j],
          dp[i+1][j-W[i]] + V[i]
        );
      }
    }
  }
}