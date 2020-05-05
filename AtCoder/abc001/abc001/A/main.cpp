#include <bits/stdc++.h>
using namespace std;


void solve(std::vector<long long> H){
  cout << H[0] - H[1] << endl;
}

int main(){
    std::vector<long long> H(2);
    for(int i = 0 ; i < 2 ; i++){
        scanf("%lld",&H[i]);
    }
    solve(std::move(H));
    return 0;
}

