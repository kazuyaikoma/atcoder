#include <bits/stdc++.h>
using namespace std;

void solve(long long N, std::vector<long long> H){
  long long cnt = 0;
  unordered_map<long, long> m1, m2;
  // I < J, I = i + H[i], J = j - H[j]を満たすi, jを探せば良い
  for (long long i=0; i<N; i++) {
    // m1, m2の2つのmapを使い、
    // 例えばJにおいてインクリメントしたマップの値にIでアクセスした結果0より大きければ、
    // それがcntに加算される
    cnt += m1[i+H[i]];
    cnt += m2[i-H[i]];
    // それぞれ、値へのアクセス時とは逆のインデックスによりインクリメント
    // これにより、二者のペアがマッチした状態のmapの値に、上の二行でアクセスできる
    m2[i+H[i]]++;
    m1[i-H[i]]++;
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
