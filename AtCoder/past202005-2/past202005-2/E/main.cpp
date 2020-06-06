#include <bits/stdc++.h>
using namespace std;


void solve(long long N, long long M, long long Q, vector<long long> x, vector<long long> y, vector<long long> c, vector<vector<long long>> A) {
  map<long long, pair<vector<long long>, long long>> m;
  for (int i=0; i<N; i++) {
    m[i] = make_pair(vector<long long>(), c[i]);
  }

  for (int i=0; i<M; i++) {
    m[x[i]-1].first.push_back(y[i]-1);
    m[y[i]-1].first.push_back(x[i]-1);
  }

  for (vector<long long>& a : A) {
    cout << m[a[1]-1].second << endl;
    if (a[0] == 2) {
      m[a[1]-1].second = a[2];
    } else {
      for (long long aj : m[a[1]-1].first) {
        m[aj].second = m[a[1]-1].second;
      }
    }
  }
}

int main(){
  long long N;
  scanf("%lld",&N);
  long long M;
  scanf("%lld",&M);
  long long Q;
  scanf("%lld",&Q);

  std::vector<long long> x(M);
  std::vector<long long> y(M);
  for (int i = 0 ; i < M ; i++){
      scanf("%lld",&x[i]);
      scanf("%lld",&y[i]);
  }

  std::vector<long long> c(N);
  for (int i = 0 ; i < N ; i++){
      scanf("%lld",&c[i]);
  }

  vector<vector<long long>> A = {};
  for(int i = 0 ; i < Q ; i++){
    vector<long long> vec = {};
    long long indicator;
    long long a;
    long long b;
    scanf("%lld",&indicator);
    if (indicator == 1) {
      cin >> a;
      vec = {1, a};
    } else {
      cin >> a;
      cin >> b;
      vec = {2, a, b};
    }
    A.push_back(vec);
  }

  solve(N, M, Q, std::move(x), std::move(y), std::move(c), std::move(A));
  return 0;
}
