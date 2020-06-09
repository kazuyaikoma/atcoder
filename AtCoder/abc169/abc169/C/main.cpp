#include <bits/stdc++.h>
using namespace std;

vector<string> split(string str, char del) {
  int first = 0;
  int last = str.find_first_of(del);
  vector<string> result;

  while (first < str.size()) {
    result.push_back(str.substr(first, last - first));
    first = last + 1;
    last = str.find_first_of(del, first);
    if (last == string::npos) last = str.size();
  }
  return result;
}

void solve(long long A, long double B){
  vector<string> v = split(to_string(B), '.');
  long long b = stol(v[0] + v[1][0] + v[1][1]);
  long long ret = A * b / 100;
  cout << ret << endl;
}

int main(){
    long long A;
    scanf("%lld",&A);
    long double B;
    scanf("%Lf",&B);
    solve(A, B);
    return 0;
}
