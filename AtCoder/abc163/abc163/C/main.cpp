#include <bits/stdc++.h>
using namespace std;


void solve(long long N, std::vector<long long> A){
  vector<int> ps = {};
  for (int i=0; i<N; i++) ps.push_back(0);

  for (int i=0; i<N-1; i++) ps[A[i]-1]++;
  for (int i=0; i<N; i++) {
    cout << ps[i] << endl;
  }
}

// Generated by 1.1.6 https://github.com/kyuridenamida/atcoder-tools  (tips: You use the default template now. You can remove this line by using your custom template)
int main(){
    long long N;
    scanf("%lld",&N);
    std::vector<long long> A(N-2+1);
    for(int i = 0 ; i < N-2+1 ; i++){
        scanf("%lld",&A[i]);
    }
    solve(N, std::move(A));
    return 0;
}
