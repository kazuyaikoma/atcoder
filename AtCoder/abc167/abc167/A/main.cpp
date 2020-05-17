#include <bits/stdc++.h>
using namespace std;

const string YES = "Yes";
const string NO = "No";

void solve(std::string S, std::string T){
  string ans = (S == T.substr(0, T.size()-1)) ? YES : NO;
  cout << ans << endl;
}

int main(){
    std::string S;
    std::cin >> S;
    std::string T;
    std::cin >> T;
    solve(S, T);
    return 0;
}
