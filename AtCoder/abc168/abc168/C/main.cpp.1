#include <bits/stdc++.h>
#define all(v) (v).begin(),(v).end()
using namespace std;


bool check(vector<long long> scores, long long i) {
  auto iter = min_element(all(scores));
  return (*iter >= i);
}

void solve(long long N, long long M, long long X, vector<vector<long long>> A){
  long long total = 0;
  for (vector<long long> m : A) total += m[0];
  long long original = total;


  do {
    vector<long long> scores(M);
    fill(all(scores), 0);
    long long cur_total = 0;

    for (int i=0; i<N; i++) {
      if (A[i][0] > total) break;
      cur_total += A[i][0];
      for (int m=0; m<M; m++) scores[m] += A[i][m+1];
      if (check(scores, X)) break;
    }
    total = min(total, cur_total);
  } while (next_permutation(all(A)));

  if (total == original) total = -1;
  cout << total << endl;
}

int main(){
  long long N;
  scanf("%lld",&N);
  long long M;
  scanf("%lld",&M);
  long long X;
  scanf("%lld",&X);
  vector<vector<long long>> A = {};
  for (int i=0; i<N; i++) {
    vector<long long> tmp(M+1);
    for(int j=0 ; j<=M; j++){
      scanf("%lld",&tmp[j]);
    }
    A.push_back(tmp);
  }
  solve(N, M, X, move(A));
  return 0;
}
