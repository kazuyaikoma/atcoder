#include <bits/stdc++.h>
using namespace std;


void solve(std::string T){
  for (int i=0; i<T.size(); i++) if (T[i] == '?') T[i] = 'D';
  cout << T << endl;
}

int main(){
    std::string T;
    std::cin >> T;
    solve(T);
    return 0;
}
