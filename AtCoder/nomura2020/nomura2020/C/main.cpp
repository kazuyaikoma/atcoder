#include <bits/stdc++.h>
using namespace std;


void solve(long long N, std::vector<long long> A){
  int ret = 0;
  vector<int> slots(N, 0);

  int slot = 1;
  for (int i=1; i<=N; i++) {
    slot = slot * 2 - A[i];
    if (slot < 0) {
      ret = -1;
      break;
    }
    slots.push_back(slot);
  }


  for (int i=0; i<N; i++) slots[i] = pow(2, i+1);
  for (int i=0; i<N; i++) {
    if (A[i+1] > slots[i]) ret = -1;
  }

  if (ret < 0) {
    cout << -1 << endl;
    return;
  }

  int idx = 1;
  for (int i=N; i>=1; i--) {
    if (A[i] < 1) continue;

    bool invalid = true;
    for (int j=idx; j<=i; j++) {
      slots[j] -= A[i];
      if (slots[j] < 0) idx++;
      else invalid = false;
    }
    if (invalid) {
      ret = -1;
      break;
    }
    ret += A[i] * i;
  }

  cout << ret << endl;
}

int main(){
    long long N;
    scanf("%lld",&N);
    std::vector<long long> A(N-0+1);
    for(int i = 0 ; i < N-0+1 ; i++){
        scanf("%lld",&A[i]);
    }
    solve(N, std::move(A));
    return 0;
}
