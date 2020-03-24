#include <bits/stdc++.h>
using namespace std;

bool isPrime(long long X) {
  if (X < 2) return false;
  if (X == 2) return true;
  if (X % 2 == 0) return false;

  int sq = sqrt(X);
  for (int i=3; i<=sq; i+=2) if (X % i == 0) return false;
  return true;
}

void solve(long long X){
    while(!isPrime(X)) X++;
    cout << X << endl;
}

// Generated by 1.1.6 https://github.com/kyuridenamida/atcoder-tools  (tips: You use the default template now. You can remove this line by using your custom template)
int main(){
    long long X;
    scanf("%lld",&X);
    solve(X);
    return 0;
}
