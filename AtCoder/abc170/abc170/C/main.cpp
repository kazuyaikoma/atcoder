#include <bits/stdc++.h>
using namespace std;

int find(vector<long long>& v, long long val) {
  for (int i=0; i<v.size(); i++) if (v[i] == val) return i;
  return -1;
}

void solve(long long X, long long N, std::vector<long long> p){
  sort(p.begin(), p.end());
  int idx = find(p, X);
  if (idx < 0) {
    cout << X << endl;
    return;
  }

  int l = idx, r = idx;
  int l_val = X, r_val = X;
  while (true) {
    if (l-1 < 0) {
      cout << l_val - 1 << endl;
      return;
    }
    if (p[--l] != --l_val) {
      cout << l_val << endl;
      return;
    }

    if (p.size()-1 < r+1) {
      cout << r_val+1 << endl;
      return;
    }
    if (p[++r] != ++r_val) {
      cout << r_val << endl;
      return;
    }

  }

}

int main(){
    long long X;
    scanf("%lld",&X);
    long long N;
    scanf("%lld",&N);
    std::vector<long long> p(N);
    for(int i = 0 ; i < N ; i++){
        scanf("%lld",&p[i]);
    }
    solve(X, N, std::move(p));
    return 0;
}
