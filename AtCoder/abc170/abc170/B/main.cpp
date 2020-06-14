#include <bits/stdc++.h>
using namespace std;

const string YES = "Yes";
const string NO = "No";

void solve(long long X, long long Y){
  for (int i=0; i<=X; i++) {
    if (i*2 + (X-i)*4 == Y) {
      cout << YES << endl;
      return;
    }
  }
  cout << NO << endl;
}

int main(){
    long long X;
    scanf("%lld",&X);
    long long Y;
    scanf("%lld",&Y);
    solve(X, Y);
    return 0;
}
