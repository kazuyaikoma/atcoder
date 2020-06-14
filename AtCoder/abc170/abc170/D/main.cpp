#include <bits/stdc++.h>
using namespace std;


void solve(long long N, std::vector<long long> A){
  sort(A.begin(), A.end());
  for (long long idx = 0; idx < A.size(); idx++) {
    auto iter = A.begin()+idx+1;
    while (iter != A.end()) {
      if (*iter % A[idx] == 0) {
        iter = A.erase(iter);
      }
      else iter++;
    }
  }

  long long ret = 0;
  if (A.size() > 1) ret = A.size();
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
