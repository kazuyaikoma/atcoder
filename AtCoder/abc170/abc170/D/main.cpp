#include <bits/stdc++.h>
using namespace std;


void solve(long long N, std::vector<long long> A){
  const long long M = 1'000'000;
  vector<long long> cnt(M, 0);
  for (long long a : A) {
    if (cnt[a] != 0) {
      cnt[a] = 2;
      continue;
    }
    for (long long i=a; i<=M; i+=a) cnt[i]++;
  }

  long long ans = 0;
  for (auto a : A) {
    if (cnt[a] == 1) ans++;
  }
  cout << ans << endl;
}

int main(){
    long long N;
    scanf("%lld",&N);
    std::vector<long long> A(N);
    for(int i = 0 ; i < N ; i++){
        scanf("%lld",&A[i]);
    }
    solve(N, std::move(A));
    return 0;
}
