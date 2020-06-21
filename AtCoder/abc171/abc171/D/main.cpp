#include <bits/stdc++.h>
using namespace std;


void solve(long long N, std::vector<long long> A, long long Q, std::vector<long long> B, std::vector<long long> C){
  map<long long, long long> m;
  for (long long a : A) {
    if (m.count(a)) m[a] += 1;
    else m[a] = 1;
  }

  long long prev = -1;
  for (long long q=0; q<Q; q++) {
    long long minus = B[q] * m[B[q]];
    long long plus = C[q] * m[B[q]];
    m[C[q]] += m[B[q]];
    m[B[q]] = 0;
    if (prev != -1) {
      cout << prev - minus + plus << endl;
      prev = prev - minus + plus;
    } else {
      long long sum = 0;
      for (auto p : m) sum += p.first * p.second;
      cout << sum << endl;
      prev = sum;
    }
  }
}

int main(){
    long long N;
    scanf("%lld",&N);
    std::vector<long long> A(N);
    for(int i = 0 ; i < N ; i++){
        scanf("%lld",&A[i]);
    }
    long long Q;
    scanf("%lld",&Q);
    std::vector<long long> B(Q);
    std::vector<long long> C(Q);
    for(int i = 0 ; i < Q ; i++){
        scanf("%lld",&B[i]);
        scanf("%lld",&C[i]);
    }
    solve(N, std::move(A), Q, std::move(B), std::move(C));
    return 0;
}
