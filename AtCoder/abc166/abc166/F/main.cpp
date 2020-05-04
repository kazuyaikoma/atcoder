#include <bits/stdc++.h>
using namespace std;

const string YES = "Yes";
const string NO = "No";

long long n;
vector<string> S;
string ans[100005];


void dfs(long long i, long long A, long long B, long long C) {
  if (A < 0 || B < 0 || C < 0) return;
  if (i == n) {
    cout << "Yes" << endl;
    for (int j=0; j<n; j++) cout << ans[j] << endl;
    exit(0);
  } else {
    if (S[i] == "AB") {
      ans[i] = 'A';
      dfs(i+1, A+1, B-1, C);
      ans[i] = 'B';
      dfs(i+1, A-1, B+1, C);
    } else if (S[i] == "AC") {
      ans[i] = 'A';
      dfs(i+1, A+1, B, C-1);
      ans[i] = 'C';
      dfs(i+1, A-1, B, C+1);
    } else {
      ans[i] = 'B';
      dfs(i+1, A, B+1, C-1);
      ans[i] = 'C';
      dfs(i+1, A, B-1, C+1);
    }
  }
}

void solve(long long N, long long A, long long B, long long C, std::vector<std::string> s){
  n = N;
  S = s;
  dfs(0, A, B, C);
  cout << "No" << endl;
}

int main(){
    long long N;
    scanf("%lld",&N);
    long long A;
    scanf("%lld",&A);
    long long B;
    scanf("%lld",&B);
    long long C;
    scanf("%lld",&C);
    std::vector<std::string> s(N);
    for(int i = 0 ; i < N ; i++){
        std::cin >> s[i];
    }
    solve(N, A, B, C, std::move(s));
    return 0;
}
