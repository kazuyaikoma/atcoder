#include <bits/stdc++.h>
using namespace std;


void solve(long long A, long long B){
  if (A <= 12) B /= 2;
  if (A <= 5) B = 0;
  cout << B << endl;
}

int main(){
  long long A;
  scanf("%lld",&A);
  long long B;
  scanf("%lld",&B);
  solve(A, B);
  return 0;
}
