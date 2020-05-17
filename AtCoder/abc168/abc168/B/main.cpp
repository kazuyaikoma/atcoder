#include <bits/stdc++.h>
using namespace std;


void solve(long long K, std::string S){
  if (S.size() > K) {
    S = S.substr(0, K);
    S += "...";
  }
  cout << S << endl;
}

int main(){
    long long K;
    scanf("%lld",&K);
    std::string S;
    std::cin >> S;
    solve(K, S);
    return 0;
}
