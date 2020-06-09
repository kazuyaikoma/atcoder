#include <bits/stdc++.h>
using namespace std;


void solve(long long N, std::vector<long long> A){
  long long max_num = 1'000'000'000'000'000'000;
  sort(A.begin(), A.end());

  if (A.size() < 1 || A[0] == 0) {
    cout << 0 << endl;
    return;
  }

  unsigned long long ret = 1;
  for (int i=0; i<N; i++) {
    long long div = max_num / ret;
    if (div + 1 < A[i]) {
      cout << -1 << endl;
      return;
    }
    ret *= A[i];
    if (ret > max_num) {
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
