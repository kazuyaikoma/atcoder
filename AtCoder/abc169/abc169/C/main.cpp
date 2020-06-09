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
  long double ret = A * B * 100;
  ret /= 100;
  vector<string> v = split(to_string(ret), '.');
  cout << v[0] << endl;
}

int main(){
    long long A;
    scanf("%lld",&A);
    long double B;
    scanf("%Lf",&B);
    solve(A, B);
    return 0;
}
