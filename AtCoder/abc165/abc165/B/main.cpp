#include <bits/stdc++.h>
using namespace std;


void solve(long long X){
  long long n = 0;
  long long bank = 100;
  while (bank < X) {
    bank *= 1.01;
    n++;
  }
  cout << n << endl;
}

// Generated by 1.1.6 https://github.com/kyuridenamida/atcoder-tools  (tips: You use the default template now. You can remove this line by using your custom template)
int main(){
    long long X;
    scanf("%lld",&X);
    solve(X);
    return 0;
}
