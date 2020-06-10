#include <bits/stdc++.h>
using namespace std;


void solve(long long N, long long M, std::vector<long long> L, std::vector<long long> R){
  long long l = L[0], r = R[0];
  for (int i=1; i<M; i++) {
    if (R[i] < l || r < L[i]) {
      cout << 0 << endl;
      return;
    }
    if (l < L[i]) l = L[i];
    if (R[i] < r) r = R[i];
    if (r < l) {
      cout << 0 << endl;
      return;
    }
  }
  cout << r - l + 1 << endl;
}

int main(){
    long long N;
    scanf("%lld",&N);
    long long M;
    scanf("%lld",&M);
    std::vector<long long> L(M);
    std::vector<long long> R(M);
    for(int i = 0 ; i < M ; i++){
        scanf("%lld",&L[i]);
        scanf("%lld",&R[i]);
    }
    solve(N, M, std::move(L), std::move(R));
    return 0;
}
