#include <bits/stdc++.h>
using namespace std;


void convert(string& str) {
  for (unsigned int i=0; i<str.size(); i++) if (i == 0 || i == str.size()-1) str[i] = str[i] - 32;
}

void solve(std::string S){
  unsigned int i = 1;
  unsigned int j = 1;
  vector<string> vec = {};
  while (i < S.size()) {
    j = i;
    while (islower(S[i])) i++;
    string s = S.substr(j-1, i-j+2);
    transform(s.begin(), s.end(), s.begin(), ::tolower);
    vec.push_back(s);
    i += 2;
  }
  sort(vec.begin(), vec.end());
  for (unsigned int i=0; i<vec.size(); i++) convert(vec[i]);
  string ans = "";
  for (string s : vec) ans += s;
  cout << ans << endl;
}

int main(){
  std::string S;
  std::cin >> S;
  solve(S);
  return 0;
}
