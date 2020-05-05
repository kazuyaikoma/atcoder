#include <bits/stdc++.h>
using namespace std;

const long long MOD = 1000000007;

void solve(long long N, long long K){
  long long ans = 0;
  cout << (1 << 10) << endl;
  for (long long bit=1; bit <= (1 << (N + 1)); bit++) {
    cout << bit << endl;
    if (std::bitset<1000000>(bit).count() >= K) ans++;
  }
  ans %= MOD;
  cout << ans - 1 << endl;
}

// Generated by 1.1.6 https://github.com/kyuridenamida/atcoder-tools  (tips: You use the default template now. You can remove this line by using your custom template)
int main(){
    long long N;
    scanf("%lld",&N);
    long long K;
    scanf("%lld",&K);
    solve(N, K);
    return 0;
}