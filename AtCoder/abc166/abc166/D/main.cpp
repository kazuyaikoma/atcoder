#include <bits/stdc++.h>
using namespace std;


void solve(long long X){
  bool solved = false;
  for (int a=-1000; a<1000; a++) {
    if (solved) return;
    for (int b=-1000; b<1000; b++) {
      if (pow(a, 5) - pow(b, 5) == X) {
        cout << a << " " << b << endl;
        solved = true;
        return;
      }
    }
  }
}

int main(){
    long long X;
    scanf("%lld",&X);
    solve(X);
    return 0;
}
