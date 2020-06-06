#include <bits/stdc++.h>
#include <iostream>
#include <cstdlib>

#include <algorithm>
#include <string>
using namespace std;


// 全て小文字に変換
void convert(string& s) {
  for (char& c : s) {
    if (64 < c && c < 91) c += 32;
  }
}

void solve(std::string s, std::string t){
  string ret = "same";
  if (s != t) {
    convert(s);
    convert(t);
    ret = (s == t) ? "case-insensitive" : "different";
  }
  cout << ret << endl;
}

int main(){
    std::string s;
    std::cin >> s;
    std::string t;
    std::cin >> t;
    solve(s, t);
    return 0;
}
