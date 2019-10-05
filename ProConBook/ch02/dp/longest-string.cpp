#include <iostream>
using namespace std;

const int N = 4;
const int M = 4;
char S[N+1] = "abcd";
char T[N+1] = "becd";

int dp[N+1][N+1];

void calc(void);

int main() {
  calc();
  printf("%d\n", dp[M][M]);
  return 0;
}

void calc() {
  for (int i=0; i<N; i++) {
    for (int j=0; j<N; j++) {
      if (S[i+1] == T[j+1]) {
        dp[i+1][j+1] = dp[i][j] + 1;
      } else {
        dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1]);
      }
    }
  }
}