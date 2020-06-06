#include <bits/stdc++.h>
using namespace std;


void solve(long long N, std::vector<std::string> A){
  long long i = 0, j = N-1;
  if (N == 1) {
    cout << A[0][0] << endl;
    return;
  }

  string a = "", b = "";

  while (i < j) {
    bool result = false;

    for (char ci : A[i]) {
      bool flag = false;
      for (char cj : A[j]) {
        if (ci == cj) {
          a.push_back(ci);
          b.push_back(cj);
          flag = true;
          break;
        }
      }
      if (flag) {
        result = true;
        break;
      }
    }

    if (!result) {
      cout << -1 << endl;
      return;
    }
    i++;
    j--;
  }

  reverse(b.begin(), b.end());
  string ret = (N % 2 != 0) ? a + A[N/2][0] + b : a + b;
  cout << ret << endl;
}

int main(){
    long long N;
    scanf("%lld",&N);
    std::vector<std::string> a(N);
    for(int i = 0 ; i < N ; i++){
        std::cin >> a[i];
    }
    solve(N, std::move(a));
    return 0;
}
