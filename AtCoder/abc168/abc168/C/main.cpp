#include <bits/stdc++.h>
using namespace std;


void solve(long long A, long long B, long long H, long long M){
  static const double pi = 3.141592653589793;
  double hm = M;
  hm /= 12;
  double tmp = abs(5*H+hm-M);
  double c = (tmp >= 30) ? 60-tmp : tmp;
  c = 2*A*B * cos(c / 30 * pi);
  double r = pow(A, 2) + pow(B, 2) - c;
  double ans = sqrt(r);
  cout << fixed << setprecision(10) << ans << endl;
}

int main(){
    long long A;
    scanf("%lld",&A);
    long long B;
    scanf("%lld",&B);
    long long H;
    scanf("%lld",&H);
    long long M;
    scanf("%lld",&M);
    solve(A, B, H, M);
    return 0;
}
