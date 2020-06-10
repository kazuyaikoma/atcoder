#include <bits/stdc++.h>
using namespace std;


void solve(long long N, long long M, std::vector<long long> A, std::vector<long long> B, std::vector<long long> C){
  map<long long, long long> m;
  for (int i=0; i<M; i++) {
    if (m.count(C[i])) m[C[i]] += B[i];
    else m[C[i]] = B[i];
  }
  for (auto a : A) {
    if (m.count(a)) m[a]++;
    else m[a] = 1;
  }

  long long ret = 0;
  int n = 0;
  for (auto iter=m.rbegin(); iter != m.rend() && n < N; iter++) {
    for (int i=0; i<(*iter).second && n < N; i++, n++) ret += (*iter).first;
  }
  cout << ret << endl;
}

int main(){
  long long N;
  scanf("%lld",&N);
  long long M;
  scanf("%lld",&M);
  std::vector<long long> A(N);
  for(int i = 0 ; i < N ; i++){
    scanf("%lld",&A[i]);
  }
  std::vector<long long> B(M);
  std::vector<long long> C(M);
  for(int i = 0 ; i < M ; i++){
    scanf("%lld",&B[i]);
    scanf("%lld",&C[i]);
  }
  solve(N, M, std::move(A), std::move(B), std::move(C));
  return 0;
}
