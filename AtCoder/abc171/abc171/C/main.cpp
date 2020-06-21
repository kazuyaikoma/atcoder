#include <bits/stdc++.h>
using namespace std;


void solve(unsigned long long N){
  vector<int> nums = {};
  while (N >= 0) {
    if (N >= 26) {
      int r = N % 26;
      long long div = (r == 0) ? (N - r - 1) : (N - r);
      N = div / 26;
      if (r == 0) r = 26;
      nums.push_back(r);
  } else {
      if (N == 0) N = 26;
      nums.push_back(N);
      break;
    }
  }
  string s = "";
  for (int num : nums) s += 'a' + num - 1;
  reverse(s.begin(), s.end());
  if (s[0] == 'z') cout << s.substr(1, s.size()) << endl;
  else cout << s << endl;
}

int main(){
    unsigned long long N;
    scanf("%llu",&N);
    solve(N);
    return 0;
}
