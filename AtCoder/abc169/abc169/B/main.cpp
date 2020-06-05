#include <bits/stdc++.h>
using namespace std;


void solve(long long N, std::vector<long long> A){
  long long ret = 1;
  bool zero = false;
  for (int i=0; i<N; i++) {
    if (A[i] == 0) zero = true;
  }
  if (zero) {
    cout << 0 << endl;
    return;
  }

  for (int i=0; i<N; i++) {
    if (to_string(ret).size() + to_string(A[i]).size() - 1 >= 20) {
      cout << -1 << endl;
      return;
    }
    ret *= A[i];
    if (ret > 1000000000000000000) {
      cout << -1 << endl;
      return;
    };
  }
  cout << ret << endl;
}

int main(){
    long long N;
    scanf("%lld",&N);
    std::vector<long long> A(N);
    for(int i = 0 ; i < N ; i++){
        scanf("%lld",&A[i]);
    }
    solve(N, std::move(A));
    return 0;
}
