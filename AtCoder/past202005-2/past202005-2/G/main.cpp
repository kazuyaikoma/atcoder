#include <bits/stdc++.h>
using namespace std;

void solve(long long N, long long X, long long Y, std::vector<long long> x, std::vector<long long> y){
  vector<vector<long long>> maze(401, vector<long long>(401, INT_MAX));
  maze[201][201] = 0;
  maze[201+X][201+Y] = 0;

  queue<pair<vector<long long>, long long>> sq, gq;
  vector<long long> s = {1, 1}, g = {X, Y};
  sq.push(make_pair(s, 0));
  gq.push(make_pair(g, 0));

  for (int i=0; i<N; i++) {
    maze[201+x[i]][201+y[i]] = -1;
  }

  vector<vector<long long>> dirs;
  vector<long long> tmp1 = {-1, 1};
  vector<long long> tmp2 = {0, 1};
  vector<long long> tmp3 = {1, 1};
  vector<long long> tmp4 = {-1, 0};
  vector<long long> tmp5 = {0, -1};
  vector<long long> tmp6 = {1, 0};
  dirs = {tmp1, tmp2, tmp3, tmp4, tmp5, tmp6};

  vector<vector<long long>> dirs2;
  vector<long long> ttmp1 = {-1, -1};
  vector<long long> ttmp2 = {0, 1};
  vector<long long> ttmp3 = {1, -1};
  vector<long long> ttmp4 = {-1, 0};
  vector<long long> ttmp5 = {0, -1};
  vector<long long> ttmp6 = {1, 0};
  dirs2 = {ttmp1, ttmp2, ttmp3, ttmp4, ttmp5, ttmp6};

  while (!sq.empty() || !gq.empty()) {
    auto pair_s = sq.front();
    vector<long long>& cur_s = pair_s.first;
    sq.pop();
    auto pair_g = gq.front();
    vector<long long>& cur_g = pair_g.first;
    gq.pop();

    for (auto& d : dirs) {
      vector<long long> ns = {cur_s[0] + d[0], cur_s[1] + d[1]};
      if (find(obj, ns)) continue;

    }
  }
}

int main(){
    long long N;
    scanf("%lld",&N);
    long long X;
    scanf("%lld",&X);
    long long Y;
    scanf("%lld",&Y);
    std::vector<long long> x(N);
    std::vector<long long> y(N);
    for(int i = 0 ; i < N ; i++){
        scanf("%lld",&x[i]);
        scanf("%lld",&y[i]);
    }
    solve(N, X, Y, std::move(x), std::move(y));
    return 0;
}
