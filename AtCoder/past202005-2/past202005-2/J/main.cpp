#include <bits/stdc++.h>
using namespace std;

// 条件を満たすkeyを二分探索してkeyを返す
long long find(map<long long, long long>& m, long long val, long long N) {
  if (m[0] < val) return 0;
  if (m[N-1] >= val) return -1;

  long long l = 0, r = N-1;
  while (l != r) {
    long long mid = (l+r)/2;
    if (m[mid] >= val) l = mid+1;
    else r = mid;
  }
  return r;
}

void solve(long long N, long long M, std::vector<long long> a){
  map<long long, long long> m;
  for (long long i=1; i<=N; i++) m[i] = 0;
  for (long long num : a) {
    long long key = find(m, num, N);
    if (key > -1) cout << key + 1 << endl;
    else cout << key << endl;
    m[key] = num;
  }
}

int main(){
    long long N;
    scanf("%lld",&N);
    long long M;
    scanf("%lld",&M);
    std::vector<long long> a(M);
    for(int i = 0 ; i < M ; i++){
        scanf("%lld",&a[i]);
    }
    solve(N, M, std::move(a));
    return 0;
}
