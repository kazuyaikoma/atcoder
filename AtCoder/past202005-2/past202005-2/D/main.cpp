#include <bits/stdc++.h>
using namespace std;


void solve(long long N, std::vector<std::string> s){
  string nums = "";
  for (int i=1; i<4*N+1; i+=4) {
    vector<string> tmps = {};
    for (int j=0; j<5; j++) {
      string tmp = string(s[j].begin()+i, s[j].begin()+i+3);
      tmps.push_back(tmp);
    }
    if (tmps[0] == ".#.") nums.push_back('1');
    else if (tmps[0] == "#.#") nums.push_back('4');
    else if (tmps[4] == "..#") nums.push_back('7');
    else if (tmps[2] == "#.#") nums.push_back('0');
    else if (tmps[1] == "..#") {
      if (tmps[3] == "#..") nums.push_back('2');
      else nums.push_back('3');
    }
    else if (tmps[1] == "#..") {
      if (tmps[3] == "..#") nums.push_back('5');
      else nums.push_back('6');
    }
    else {
      if (tmps[3] == "..#") nums.push_back('9');
      else nums.push_back('8');
    }
  }
  cout << nums << endl;
}

int main(){
    long long N;
    scanf("%lld",&N);
    std::vector<std::string> s(5);
    for(int i = 0 ; i < 5 ; i++){
        std::cin >> s[i];
    }
    solve(N, std::move(s));
    return 0;
}
