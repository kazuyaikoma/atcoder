#include <bits/stdc++.h>
using namespace std;


void solve(long long A, long long R, long long N){
  if (N == 1) {
    cout << A << endl;
    return;
  }
  long long m = 1'000'000'000;
  long long powed = R;
  for (long long i=0; i<N-2; i++) {
    powed *= R;
    if (powed > m) {
      cout << "large" << endl;
      return;
    }
  }

  if (powed * A > m) {
    cout << "large" << endl;
    return;
  }
  cout << powed * A << endl;
}

int main(){
    long long A;
    scanf("%lld",&A);
    long long R;
    scanf("%lld",&R);
    long long N;
    scanf("%lld",&N);
    solve(A, R, N);
    return 0;
}
