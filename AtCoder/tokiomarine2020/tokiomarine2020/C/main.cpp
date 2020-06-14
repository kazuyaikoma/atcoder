#include <bits/stdc++.h>
#define rep(i, n) for (int i=0; i<n; ++i)
using namespace std;
using ll = long long;


void solve(ll N, ll K, std::vector<ll> A){
  vector<ll> v(N, 0);
  rep(i, K) {
    vector<ll> v(N+1);
    rep(j, N) {
      int l = max(0, j-(int)A[j]);
      int r = min((int)N, j+(int)A[j]+1);
      v[l]++;
      v[r]--;
    }

    rep(j, N) v[j+1] += v[j];
    v.pop_back();
    if (A == v) break;
    A = v;
  }

  rep(i, N) cout << A[i] << endl;
}

int main(){
  ll N;
  scanf("%lld",&N);
  ll K;
  scanf("%lld",&K);
  std::vector<ll> A(N);
  for(int i = 0 ; i < N ; i++){
    scanf("%lld",&A[i]);
  }
  solve(N, K, std::move(A));
  return 0;
}
