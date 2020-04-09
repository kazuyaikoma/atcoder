// TODO: this code is somehing wrong.
#include <bits/stdc++.h>
#define ALL(v) (v).begin(),(v).end()
using namespace std;
using ll = long long;


void solve(ll N, ll Q, vector<vector<ll>> A){
  vector<string> ans = {};
  for (int i=0; i<N; i++) {
    string s = "";
    for (int j=0; j<N; j++) s += 'N';
    ans.push_back(s);
  }

  for (vector<ll> q : A) {
    if (q[0] == 1) ans[q[1]-1][q[2]-1] = 'Y';
    if (q[0] == 2) for (unsigned int i=0; i<N; i++) if (ans[i][q[1]-1] == 'Y') ans[q[1]-1][i] = 'Y';
    if (q[0] == 3) {
      vector<int> follows = {};
      for (unsigned int i=0; i<N; i++) if (ans[q[1]-1][i] == 'Y') follows.push_back(i);
      for (int i : follows) for (unsigned int j=0; j<N; j++) if (ans[i][j] == 'Y') ans[q[1]-1][j] = 'Y';
    }
  }

  for (string s : ans) cout << s << endl;
}

int main(){
  ll N;
  scanf("%lld", &N);
  ll Q;
  scanf("%lld", &Q);

  vector<vector<ll>> A = {};
  for(int i = 0; i < Q; i++){
    ll mark;
    scanf("%lld", &mark);
    if (mark == 1) {
      ll first;
      ll second;
      scanf("%lld", &first);
      scanf("%lld", &second);
      A.insert(A.end(), {mark, first, second});
    } else {
      ll first;
      scanf("%lld", &first);
      A.insert(A.end(), {mark, first});
    }
  }
  solve(N, Q, move(A));
  return 0;
}
