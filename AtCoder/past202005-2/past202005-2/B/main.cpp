#include <bits/stdc++.h>
using namespace std;

void solve(long long N, long long M, long long Q, std::vector<vector<long long>> A){
  vector<vector<long long>> users(N, vector<long long>());
  vector<long long> problems(M, N);
  for (vector<long long>& v : A) {
    if (v[0] == 2) {
      if (problems[v[2]-1] > 0) problems[v[2]-1]--;
      users[v[1]-1].push_back(v[2]-1);
    } else {
      long long score = 0;
      for (auto s : users[v[1]-1]) {
        score += problems[s];
      }
      cout << score << endl;
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
    vector<vector<long long>> A = {};
    for(int i = 0 ; i < Q ; i++){
      vector<long long> vec = {};
      long long indicator;
      long long x;
      long long y;
      scanf("%lld",&indicator);
      if (indicator == 1) {
        cin >> x;
        vec = {1, x};
      } else {
        cin >> x;
        cin >> y;
        vec = {2, x, y};
      }
      A.push_back(vec);
    }
    solve(N, M, Q, std::move(A));
    return 0;
}
