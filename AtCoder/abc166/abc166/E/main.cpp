#include <bits/stdc++.h>
using namespace std;

void solve(long long N, std::vector<long long> H){
  long long cnt = 0;
  for (long long i=N-1; i>=0; i--) {
    long long partner_max = i + H[i];
    if (partner_max < 0) continue;
    for (long long j=0; j<partner_max; j++) if (i-j == H[i]+H[j]) cnt++;
  }
  cout << cnt << endl;
}

int main(){
    long long N;
    scanf("%lld",&N);
    std::vector<long long> H(N);
    for(long long i = 0 ; i < N ; i++){
        scanf("%lld",&H[i]);
    }
    solve(N, std::move(H));
    return 0;
}
