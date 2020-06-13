#include <bits/stdc++.h>
using namespace std;

const string YES = "YES";
const string NO = "NO";

void solve(long long A, long long V, long long B, long long W, long long T){
  if (W >= V) {
    cout << NO << endl;
    return;
  }

  long long abs_ab = abs(A-B), abs_vw = abs(V-W);
  if (abs_ab % abs_vw != 0) {
    cout << NO << endl;
    return;
  }

  if (abs_ab / abs_vw > T) cout << NO << endl;
  else cout << YES << endl;
}

int main(){
    long long A;
    scanf("%lld",&A);
    long long V;
    scanf("%lld",&V);
    long long B;
    scanf("%lld",&B);
    long long W;
    scanf("%lld",&W);
    long long T;
    scanf("%lld",&T);
    solve(A, V, B, W, T);
    return 0;
}
