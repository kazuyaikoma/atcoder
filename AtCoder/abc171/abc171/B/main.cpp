#include <bits/stdc++.h>
using namespace std;


void solve(long long N, long long K, std::vector<long long> p){
  sort(p.begin(), p.end());
  long long sum = 0;
  for (long long i=0; i<K; i++) sum += p[i];
  cout << sum << endl;
}

int main(){
    long long N;
    scanf("%lld",&N);
    long long K;
    scanf("%lld",&K);
    std::vector<long long> p(N);
    for(int i = 0 ; i < N ; i++){
        scanf("%lld",&p[i]);
    }
    solve(N, K, std::move(p));
    return 0;
}
