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

int fixed(int time, bool first) {
  int one_digit = time % 10;
  int two_digit = time % 100;
  if (!first && 57<=two_digit) return time + (100-two_digit);

  if (first) {
    if (one_digit < 5) return time - one_digit;
    else return time - one_digit + 5;
  } else {
    if (0 < one_digit && one_digit < 5) return time + (5 - one_digit);
    else if (5 < one_digit && one_digit < 10) return time + (10 - one_digit);
  }
  return time;
}

void solve(long long N, std::vector<string> A){
  map<int, int> m;
  for (string a : A) {
    vector<string> time = {};
    time = split(a, '-');
    int begin = fixed(stoi(time[0]), true), end = fixed(stoi(time[1]), false);
    m[begin] = end;
  }

  vector<string> ans = {};
  int prev_begin = -1, prev_end = -1;
  for (auto elm : m) {
    if (prev_end < elm.first) {
      if (prev_begin != -1) {
        cout << setw(4) << setfill('0') << prev_begin;
        cout << '-';
        cout << setw(4) << setfill('0') << prev_end << endl;
      }

      prev_begin = elm.first;
      prev_end = elm.second;
      continue;
    }
    prev_end = max(prev_end, elm.second);
  }
  cout << setw(4) << setfill('0') << prev_begin;
  cout << '-';
  cout << setw(4) << setfill('0') << prev_end << endl;
}

int main(){
  long long N;
  scanf("%lld",&N);
  vector<string> A(N);
  for(int i = 0 ; i < N ; i++){
    string str;
    cin >> str;
    A[i] = str;
  }
  solve(N, std::move(A));
  return 0;
}
