#include <bits/stdc++.h>
using namespace std;


void solve(std::vector<long long> x){
  int ret = 0;
  for (int i=0; i<5; i++) if (x[i] == 0) ret = i+1;
  cout << ret << endl;
}

int main(){
    std::vector<long long> x(5);
    for(int i = 0 ; i < 5 ; i++){
        scanf("%lld",&x[i]);
    }
    solve(std::move(x));
    return 0;
}
