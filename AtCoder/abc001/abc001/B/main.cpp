#include <bits/stdc++.h>
using namespace std;

const long long MOD = 5;

void solve(long long m){
  int ans = 0;
  if (m < 100) {
    ans = 0;
  } else if (100 <= m && m <= 5000) {
    ans = m / 100;
  } else if (6000 <= m && m <= 30000) {
    ans = m / 1000 + 50;
  } else if (35000 <= m && m <= 70000) {
    ans = (m/1000-30)/5 + 80;
  } else if (70000 < m)  {
    ans = 89;
  }
  cout << setw(2) << setfill('0') << ans << endl;
}

int main(){
    long long m;
    scanf("%lld",&m);
    solve(m);
    return 0;
}

