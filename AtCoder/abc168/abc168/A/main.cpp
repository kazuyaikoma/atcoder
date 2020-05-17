#include <bits/stdc++.h>
using namespace std;


void solve(long long N){
  string ans = "hon";
  N %= 10;
  if (N == 0 || N == 1 || N == 6 || N == 8) ans = "pon";
  else if (N == 3) ans = "bon";
  cout << ans << endl;
}

int main(){
    long long N;
    scanf("%lld",&N);
    solve(N);
    return 0;
}
